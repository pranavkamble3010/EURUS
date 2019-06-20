package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"strconv"
	"time"
)

//Response can be an answer to a question, opinion in a discussion or can be a point in a debate
type Response struct {
	ObjectId        primitive.ObjectID
	InteractionId   primitive.ObjectID
	ResponseId      int32 //0 for answer, 1 for opinion, 2 for point
	ResponseType    int32
	OwnerId         primitive.ObjectID
	ResponseContent string
	DateCreated     primitive.DateTime //Date created
	DateUpdated     primitive.DateTime //Date updated
}

func init() {
	//uri = "mongodb+srv://admin:eurusAdmin_123@cluster0-oqmw3.mongodb.net/test?retryWrites=true"

}

// func getResponseCollection(context context.Context) (*mongo.Collection, error) {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

// 	if err != nil {
// 		fmt.Printf("Error connecting to DB! - %s", err.Error())
// 	}

// 	err = client.Connect(context)

// 	collection := client.Database("eurus").Collection("responses")
// 	return collection, err
//}

//extractResponses function iterates over the provided cursor and populates the list of responses
func extractResponses(ctx context.Context, cur *mongo.Cursor) []Response {
	var responses []Response
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			fmt.Printf("error occurred! Error is - %s", err.Error())
		}
		//fmt.Printf("%s", result["topic"])
		resultMap := result.Map()
		mapKey := resultMap["_id"].(primitive.ObjectID) //Key of the map element

		var resp Response

		resp.ObjectId = mapKey
		resp.InteractionId = resultMap["interactionId"].(primitive.ObjectID)
		resp.ResponseId = resultMap["responseId"].(int32)
		intrType, err := strconv.Atoi(resultMap["responseType"].(primitive.Decimal128).String())
		resp.ResponseType = int32(intrType)
		resp.OwnerId = resultMap["ownerId"].(primitive.ObjectID)
		resp.ResponseContent = resultMap["responseContent"].(string)
		if resultMap["dateCreated"] != nil {
			resp.DateCreated = resultMap["dateCreated"].(primitive.DateTime)
		}

		if resultMap["dateUpdated"] != nil {
			resp.DateUpdated = resultMap["dateUpdated"].(primitive.DateTime)
		}

		responses = append(responses, resp)

	}
	return responses
}

//InsertResponse function inserts a new Response
func InsertResponse(response Response) string {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "responses")

	respType, err := primitive.ParseDecimal128(strconv.Itoa(int(response.ResponseType)))

	result, err := collection.InsertOne(ctx, bson.M{
		"interactionId":   response.InteractionId,
		"responseType":    respType,
		"ownerId":         response.OwnerId,
		"responseContent": response.ResponseContent,
		"responseId":      rand.Intn(999999),
		"dateCreated":     time.Now(),
		"dateUpdated":     time.Now()})

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}

//UpdateResponse function updates an Interaction
func UpdateResponse(response Response) int64 {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "responses")
	objId := bson.D{{"_id", response.ObjectId}}

	doc := bson.D{
		{"$set", bson.D{{"responseContent", response.ResponseContent},
			{"dateUpdated", time.Now()},
		}},
	}

	result, err := collection.UpdateOne(ctx, objId, doc)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	return result.ModifiedCount
}

//DeleteResponse function deletes an Response
func DeleteResponse(objectID string) int64 {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "responses")

	objID, err := primitive.ObjectIDFromHex(objectID)

	filter := bson.D{{"_id", objID}}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		fmt.Printf("Error occured! The error is - %s", err.Error())
	}

	return result.DeletedCount

}

//GetResponsesByOwnerID function
func GetResponsesByOwnerID(OwnerID string) []Response {

	var responses []Response

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "responses")

	//Declare filter
	ownerId, err := primitive.ObjectIDFromHex(OwnerID)
	filter := bson.M{"ownerId": ownerId}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	responses = extractResponses(ctx, cur)
	defer cur.Close(ctx)

	return responses

}

//GetResponsesByResponseID function returns response by ID
func GetResponsesByResponseID(ResponseID string) Response {

	var responses []Response

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "responses")

	//Declare filter
	responseId, err := primitive.ObjectIDFromHex(ResponseID)
	filter := bson.M{"_id": responseId}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	responses = extractResponses(ctx, cur)
	defer cur.Close(ctx)

	return responses[0]

}

//GetResponsesByIntrID function returns response by interactionID of the interaction it belongs to
func GetResponsesByIntrID(intrID string) []Response {

	var responses []Response

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //Define context
	defer cancel()

	collection, err := getModelCollection(ctx, "responses")

	//Declare filter
	intrId, err := primitive.ObjectIDFromHex(intrID)
	filter := bson.M{"interactionId": intrId}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Printf("error occurred! Error is - %s", err.Error())
	}

	responses = extractResponses(ctx, cur)
	defer cur.Close(ctx)

	return responses

}
