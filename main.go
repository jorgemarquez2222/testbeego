package main

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

func main() {

	ctrl := &MainController{}

	web.BConfig.CopyRequestBody = true

	// we register the path / to &MainController
	// if we don't pass methodName as third param
	// web will use the default mappingMethods
	// GET http://localhost:8080  -> Get()
	// POST http://localhost:8080 -> Post()
	// ...
	web.Router("/", ctrl)

	web.Run()
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends web.Controller
type MainController struct {
	web.Controller
}

type user struct {
	Name     string                 `json:"name"`
	Password string                 `json:"password"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// address: http://localhost:8080 Post
func (ctrl *MainController) Post() {
	input := user{}

	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &input); err != nil {
		ctrl.Data["json"] = err.Error()
	}

	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}

func (ctrl *MainController) Get() {
	input := user{}

	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &input); err != nil {
		ctrl.Data["json"] = err.Error()
	}

	ctrl.Data["json"] = input
	ctrl.ServeJSON()
}
