package base

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
)

type BaseTemplateparams struct {
	LoggedIn 					bool
	DisplayLogin  				bool
	DisplayCreateAccount 		bool
	DisplayCreatAcctResponse 	bool
	SideNav	      				bool
	MainTable	  				bool
	Cards		  				bool
}

func GetBaseTemplateObj() BaseTemplateparams{
	return BaseTemplateparams{
		LoggedIn: false,
		DisplayLogin: true,
		DisplayCreateAccount: false,
		DisplayCreatAcctResponse: false,
		SideNav: false,
		MainTable: false,
		Cards: false,
	}
}

///////////////////// Base View //////////////////////
type BaseVw struct {
	App *config.AppConfig
}

var AppBaseVw *BaseVw

func init() {
	AppBaseVw = &BaseVw{
		App: nil,
	}
	gob.Register(BaseVwData{})
}

func (m *BaseVw) RegisterView(app *config.AppConfig) *BaseVw {
	log.Println("Registering AppLayoutVw...")
	AppBaseVw.App = app
	return AppBaseVw
}

func (m *BaseVw) RegisterHandler() constants.ViewHandler {
	return &BaseVw{}
}

// func (m *BaseVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("[lyoutvw] - Processing request")
// 	CreateBaseVwData().ProcessHttpRequest(w, r)
// }

// func (m *BaseVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) any{
// 	CreateBaseVwData().ProcessMBusRequest(w, r)
// 	return nil
// }

// func (m *BaseVw) HandleRequest(w http.ResponseWriter, r *http.Request, c chan any, d chan int) {
func (m *BaseVw) HandleRequest(w http.ResponseWriter, event constants.AppEvent) any{	
	fmt.Println("[basevw] - HandleRequest")
	var obj BaseVwData

	if session.SessionSvc.SessionMgr.Exists(event.Context, "basevw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Context, "basevw").(BaseVwData)
	} else {
		obj = *CreateBaseVwData()	
	}
	obj.ProcessHttpRequest(w, event)	
	
	session.SessionSvc.SessionMgr.Put(event.Context, "basevw", obj)

	return obj
}


///////////////////// Base View Data //////////////////////

type BaseVwData struct {
	Base BaseTemplateparams
	Data any
	View int
}

func CreateBaseVwData() *BaseVwData {
	return &BaseVwData{
		Base: GetBaseTemplateObj(),
		Data: nil,
	}
}

func (m *BaseVwData) ProcessHttpRequest(w http.ResponseWriter, event constants.AppEvent) *BaseVwData{
	fmt.Println("[basevw] - Processing request")
	m.View = constants.RM_HOME
	return m // TEMPORARY
}

func (m *BaseVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {

}