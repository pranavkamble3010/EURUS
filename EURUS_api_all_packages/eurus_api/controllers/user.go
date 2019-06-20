package controllers

import (
	"encoding/json"
	"eurus_api/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title AddUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [post]
func (u *UserController) AddUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = uid
	u.ServeJSON()

}

// @Title GetAll
// @Description get all User
// @Success 200 {object} models.Users
// @Failure 403 :objectId is empty
// @router / [get]
func (c *UserController) GetAll() {

}

// @Title GetUserByID
// @Description get user by uid
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) GetUserByID() {
	uid := u.GetString(":uid")
	if uid != "" {
		user := models.GetUserByID(uid)
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title Login
// @Description Autheticate the user
// @Param	uid		path 	string	true		"The uid you want to autheticate"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router / [post]
func (u *UserController) Login() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uu, err := models.AutheticateUser(user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = uu
	}

	u.Ctx.Output.Header("Access-Control-Allow-Headers", "application/json")
	u.ServeJSON()
}

// @Title LogOut
// @Description Logout the user
// @Param	uid		path 	string	true		"The uid you want to logout"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [post]
// func (u *UserController) LogOut() {
// }

// // @Title Delete
// // @Description delete the user
// // @Param	uid		path 	string	true		"The uid you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 uid is empty
// // @router /:uid [delete]
// func (u *UserController) Delete() {
// 	uid := u.GetString(":uid")
// 	models.DeleteUser(uid)
// 	u.Data["json"] = "delete success!"
// 	u.ServeJSON()
// }
