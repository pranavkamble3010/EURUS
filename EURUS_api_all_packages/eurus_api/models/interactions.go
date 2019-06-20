package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//Global declarations
var (
	Interactions map[primitive.ObjectID]bson.D //list of all interactions
)

//Interaction can be a question, discussion or debate
type Interaction struct {
	ObjectId        primitive.ObjectID
	InteractionId   int32
	InteractionType int32 //0 for question, 1 for discussion, 2 for debate
	Topic           string
	Tags            string
	OwnerId         primitive.ObjectID
	Description     string
	Responses       primitive.A        //list of objectIds of all responses to the interaction
	DateCreated     primitive.DateTime //Date created
	DateUpdated     primitive.DateTime
}

func init() {
	//Initialize global variables
	Interactions = make(map[primitive.ObjectID]bson.D)
}

//extractResults function iterates over the provided cursor and populates the list of interactions
func extractResults(ctx context.Context, cur *mongo.Cursor) []Interaction {
	var interactions []Interaction
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			fmt.Printf("error occurred! Error is - %s", err.Error())
		}
		//fmt.Printf("%s", result["topic"])
		resultMap := result.Map()
		mapKey := resultMap["_id"].(primitive.ObjectID) //Key of the map element

		var i Interaction

		i.ObjectId = mapKey
		i.InteractionId = resultMap["interactionId"].(int32)

		intrType, err := strconv.Atoi(resultMap["interactionType"].(primitive.Decimal128).String())
		i.InteractionType = int32(intrType)

		i.Topic = resultMap["topic"].(string)
		i.Tags = resultMap["tags"].(string)
		i.OwnerId = resultMap["ownerId"].(primitive.ObjectID)
		if resultMap["responses"] != nil {
			i.Responses = resultMap["responses"].(primitive.A)
		}

		i.Description = resultMap["description"].(string)
		if resultMap["dateCreated"] != nil {
			i.DateCreated = resultMap["dateCreated"].(primitive.DateTime)
		}

		if resultMap["dateUpdated"] != nil {
			i.DateUpdated = resultMap["dateUpdated"].(primitive.DateTime)
		}

		interactions = append(interactions, i)

	}
	return interactions
}

//GetAllInteractions function
func GetAllInteractions() []Interaction {

	var allInteractions []Interaction

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "interactions")
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	allInteractions = extractResults(ctx, cur)
	defer cur.Close(ctx)
	return allInteractions
}

//GetInteractionsByType function returns all interactions of type question
func GetInteractionsByType(interactionType string) []Interaction {

	var interactionsByType []Interaction

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "interactions")

	//Declare filter
	var filter bson.M

	switch interactionType {
	case "question":
		filter = bson.M{"interactionType": 0}
		break

	case "debate":
		filter = bson.M{"interactionType": 1}
		break

	case "discussion":
		filter = bson.M{"interactionType": 2}
		break

	}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	interactionsByType = extractResults(ctx, cur)
	defer cur.Close(ctx)

	return interactionsByType

}

//GetInteractionsByTags function returns all interactions containing the topic
//Input is comma separated topics' list
func GetInteractionsByTags(interactionType string) []Interaction {

	var interactionsByTopic []Interaction
	var pattern string

	pattern = strings.Replace(interactionType, ",", "|", -1)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "interactions")

	filter := bson.D{
		{"tags", primitive.Regex{Pattern: pattern, Options: "i"}},
	}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}
	defer cur.Close(ctx)

	interactionsByTopic = extractResults(ctx, cur)

	return interactionsByTopic

}

//InsertInteraction function inserts a new Interaction
func InsertInteraction(interaction Interaction) string {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "interactions")

	intrType, err := primitive.ParseDecimal128(strconv.Itoa(int(interaction.InteractionType)))

	result, err := collection.InsertOne(ctx, bson.M{
		"interactionId":   rand.Intn(99999999),
		"interactionType": intrType,
		"topic":           interaction.Topic,
		"tags":            interaction.Tags,
		"ownerId":         interaction.OwnerId,
		"responses":       interaction.Responses,
		"description":     interaction.Description,
		"dateCreated":     time.Now(),
		"dateUpdated":     time.Now()})

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}

//UpdateInteraction function updates an Interaction
func UpdateInteraction(interaction Interaction, updateResponsesFlag bool) int64 {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "interactions")

	intrType, err := primitive.ParseDecimal128(strconv.Itoa(int(interaction.InteractionType)))
	objId := bson.D{{"_id", interaction.ObjectId}}

	if updateResponsesFlag == true {

		doc := bson.D{
			{"$set", bson.D{{"topic", interaction.Topic},
				{"interactionType", intrType},
				{"description", interaction.Description},
				{"tags", interaction.Tags},
				//{"responses", interaction.Responses},
				{"dateUpdated", time.Now()},
			}},
		}

		result, err := collection.UpdateOne(ctx, objId, doc)

		if err != nil {
			fmt.Printf("error occurred! Error is - %s", err.Error())
		}
		return result.ModifiedCount
	}

	doc := bson.D{
		{"$set", bson.D{{"responses", interaction.Responses},
			{"dateUpdated", time.Now()},
		}},
	}

	result, err := collection.UpdateOne(ctx, objId, doc)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	return result.ModifiedCount

}

//DeleteInteraction function deletes an Interaction
func DeleteInteraction(objectID string) int64 {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "interactions")

	objID, err := primitive.ObjectIDFromHex(objectID)

	filter := bson.D{{"_id", objID}}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		fmt.Printf("Error occured! The error is - %s", err.Error)
	}

	return result.DeletedCount

}

//GetInteractionsByOwnerID returns interactions by OwnerID
func GetInteractionsByOwnerID(OwnerID string) []Interaction {

	var interactionsByType []Interaction

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "interactions")

	//Declare filter
	ownerId, err := primitive.ObjectIDFromHex(OwnerID)
	filter := bson.M{"ownerId": ownerId}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	interactionsByType = extractResults(ctx, cur)
	defer cur.Close(ctx)

	return interactionsByType

}
