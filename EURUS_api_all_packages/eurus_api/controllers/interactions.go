package controllers

import (
	"encoding/json"
	"eurus_api/models"
	"fmt"
	"github.com/astaxie/beego"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"gopkg.in/mgo.v2/bson"
	//"log"
)

//InteractionsController is controller to handle interactions
type InteractionsController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Interactions
// @Failure 403 :objectId is empty
// @router / [get]
func (c *InteractionsController) GetAll() {

	c.Data["json"] = models.GetAllInteractions()
	c.ServeJSON()
}

// @Title GetInteractionsByType
// @Description get all objects based on interactionType
//@Param	intrType		path 	string	true		"the interactiontype you want to get"
// @Success 200 {object} models.Interactions
// @Failure 403 :intrType is empty
// @router /:intrType [get]
func (c *InteractionsController) GetInteractionsByType() {

	intrType := c.Ctx.Input.Param(":intrType")

	c.Data["json"] = models.GetInteractionsByType(intrType)
	c.ServeJSON()

}

// @Title GetInteractionsByOwner
// @Description get all objects based on OwnerId
//@Param	ownerid		path 	string	true		"the interactiontype you want to get"
// @Success 200 {object} models.Interactions
// @Failure 403 :ownerid is empty
// @router /owner/:ownerid [get]
func (c *InteractionsController) GetInteractionsByOwner() {

	ownerId := c.Ctx.Input.Param(":ownerid")

	c.Data["json"] = models.GetInteractionsByOwnerID(ownerId)
	c.ServeJSON()

}

// @Title Create
// @Description create object
// @Param	body		body 	models.Interaction	true		"The Interaction content"
// @Success 200 {string} models.Interaction.Id
// @Failure 403 body is empty
// @router / [post]
func (c *InteractionsController) Create() {

	var intr models.Interaction
	json.Unmarshal(c.Ctx.Input.RequestBody, &intr)
	//fmt.Printf(intr.Topic + " " + intr.Description)
	res := models.InsertInteraction(intr)
	fmt.Printf("\n%s", res)

	if res != "" {
		c.Data["json"] = "success"
	} else {

		c.Data["json"] = "failure"
	}
	c.ServeJSON()
}

// @Title search
// @Description search by tags
// @Param	body		body 	models.Interaction	true		"The Interaction content"
// @Success 200 {string} models.Interaction.Id
// @Failure 403 body is empty
// @router /search [post]
func (c *InteractionsController) Search() {

	var intr models.Interaction
	json.Unmarshal(c.Ctx.Input.RequestBody, &intr)
	//fmt.Printf(intr.Topic + " " + intr.Description)
	res := models.GetInteractionsByTags(intr.Tags)
	fmt.Printf("\n%s", res)

	// if res != "" {
	// 	c.Data["json"] = "success"
	// } else {

	// 	c.Data["json"] = "failure"
	// }

	c.Data["json"] = res
	c.ServeJSON()
}

// @Title UpdateInteraction
// @Description update interactions
// @Param	body		body 	models.Interaction	true		"The Interaction content"
// @Success 200 {string} models.Interaction.Id
// @Failure 403 body is empty
// @router / [put]
func (c *InteractionsController) UpdateInteraction() {

	var intr models.Interaction
	json.Unmarshal(c.Ctx.Input.RequestBody, &intr)
	fmt.Printf(intr.Topic + " " + intr.Description)
	res := models.UpdateInteraction(intr, false)
	fmt.Printf("\n%d", res)

	if res != 0 {
		c.Data["json"] = "success"
	} else {

		c.Data["json"] = "failure"
	}
	c.ServeJSON()

}

// @Title DeleteInteraction
// @Description Delete interactions
// @Param	objid		path 	string	true		"The Interaction content"
// @Success 200 {string} models.Interaction.Id
// @Failure 403 body is empty
// @router /:objid [delete]
func (c *InteractionsController) DeleteInteraction() {

	intrType := c.Ctx.Input.Param(":objid")

	res := models.DeleteInteraction(intrType)

	if res != 0 {
		c.Data["json"] = "success"
	} else {

		c.Data["json"] = "failure"
	}
	c.ServeJSON()

}

// @Title UpdateInteraction
// @Description update interactions to add responseIDs to the responses array in the interactions
// @Param	body		body 	models.Interaction	true		"The Interaction content"
// @Success 200 {string} models.Interaction.Id
// @Failure 403 body is empty
// @router /updateResponses [put]
func (c *InteractionsController) UpdateResponsesByIntrID() {

	var intr models.Interaction
	json.Unmarshal(c.Ctx.Input.RequestBody, &intr)
	//fmt.Printf(intr.Topic + " " + intr.Description)
	res := models.UpdateInteraction(intr, true)
	fmt.Printf("\n%d", res)

	if res != 0 {
		c.Data["json"] = "success"
	} else {

		c.Data["json"] = "failure"
	}
	c.ServeJSON()
}
