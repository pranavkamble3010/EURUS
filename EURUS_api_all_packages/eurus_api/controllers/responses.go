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

//ResponseController is controller to handle interactions
type ResponseController struct {
	beego.Controller
}

// @Title GetResponseByID
// @Description get all responses based on ID
//@Param	responseid		path 	string	true		"the responseid of response you want to get"
// @Success 200 {object} models.Response
// @Failure 403 :responseid is empty
// @router /:responseid [get]
func (r *ResponseController) GetResponseByID() {

	responseId := r.Ctx.Input.Param(":responseid")
	resposneData := models.GetResponsesByResponseID(responseId)
	r.Data["json"] = resposneData
	r.ServeJSON()

}

// @Title GetResponsesByOwner
// @Description get all objects based on OwnerId
//@Param	ownerId		path 	string	true		"the responses of the owner"
// @Success 200 {object} models.Interactions
// @Failure 403 :ownerid is empty
// @router /owner/:ownerid [get]
func (r *ResponseController) GetResponsesByOwner() {

	ownerId := r.Ctx.Input.Param(":ownerid")

	r.Data["json"] = models.GetResponsesByOwnerID(ownerId)
	r.ServeJSON()

}

// @Title GetResponsesByIntrID
// @Description get all objects based on interaction object ID
//@Param	ownerId		path 	string	true		"the responses of the owner"
// @Success 200 {object} models.Interactions
// @Failure 403 :ownerid is empty
// @router /intrresp/:intrid [get]
func (r *ResponseController) GetResponsesByIntrID() {

	intrId := r.Ctx.Input.Param(":intrid")

	r.Data["json"] = models.GetResponsesByIntrID(intrId)
	r.ServeJSON()

}

// @Title Create
// @Description create object
// @Param	body		body 	models.Interaction	true		"The Interaction content"
// @Success 200 {string} models.Interaction.Id
// @Failure 403 body is empty
// @router / [post]
func (r *ResponseController) Create() {
	var resp models.Response
	json.Unmarshal(r.Ctx.Input.RequestBody, &resp)
	//fmt.Printf(intr.Topic + " " + intr.Description)
	res := models.InsertResponse(resp)
	fmt.Printf("\n%s", res)

	if res != "" {
		r.Data["json"] = "success"
	} else {

		r.Data["json"] = "failure"
	}
	r.ServeJSON()
}

// @Title UpdateResponse
// @Description update interactions
// @Param	body		body 	models.Interaction	true		"The Interaction content"
// @Success 200 {string} models.Interaction.Id
// @Failure 403 body is empty
// @router / [put]
func (r *ResponseController) UpdateResponse() {

	var resp models.Response
	json.Unmarshal(r.Ctx.Input.RequestBody, &resp)
	//fmt.Printf(intr.Topic + " " + intr.Description)
	res := models.UpdateResponse(resp)
	fmt.Printf("\n%d", res)

	if res != 0 {
		r.Data["json"] = "success"
	} else {

		r.Data["json"] = "failure"
	}
	r.ServeJSON()

}

// @Title DeleteResponse
// @Description Delete interactions
// @Param	objid 	path 	string	true		"The Interaction content"
// @Success 200 {string} models.Response.Id
// @Failure 403 body is empty
// @router /:objid [delete]
func (r *ResponseController) DeleteResponse() {

	resp := r.Ctx.Input.Param(":objid")

	res := models.DeleteResponse(resp)

	if res != 0 {
		r.Data["json"] = "success"
	} else {

		r.Data["json"] = "failure"
	}
	r.ServeJSON()

}
