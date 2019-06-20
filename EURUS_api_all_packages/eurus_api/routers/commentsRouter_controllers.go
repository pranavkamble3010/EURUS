package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "UpdateInteraction",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "GetInteractionsByType",
            Router: `/:intrType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "DeleteInteraction",
            Router: `/:objid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "GetInteractionsByOwner",
            Router: `/owner/:ownerid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "Search",
            Router: `/search`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:InteractionsController"],
        beego.ControllerComments{
            Method: "UpdateResponsesByIntrID",
            Router: `/updateResponses`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"],
        beego.ControllerComments{
            Method: "UpdateResponse",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"],
        beego.ControllerComments{
            Method: "DeleteResponse",
            Router: `/:objid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"],
        beego.ControllerComments{
            Method: "GetResponseByID",
            Router: `/:responseid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"],
        beego.ControllerComments{
            Method: "GetResponsesByIntrID",
            Router: `/intrresp/:intrid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:ResponseController"],
        beego.ControllerComments{
            Method: "GetResponsesByOwner",
            Router: `/owner/:ownerid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:UserController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:UserController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:UserController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserByID",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["eurus_api/controllers:UserController"] = append(beego.GlobalControllerRouter["eurus_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
