package layoutvw

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	//"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/services/session"	
	b "dshusdock/go_project/internal/views/base"
	
	"encoding/gob"

	// "dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log"
	"net/http"
	// "net/url"
)

type LayoutVw struct {
	App *config.AppConfig
}

var AppLayoutVw *LayoutVw

func init() {
	AppLayoutVw = &LayoutVw{
		App: nil,
	}
	gob.Register(LayoutVwData{})
	//messagebus.GetBus().Subscribe("Event:ViewChange", AppLayoutVw.HandleMBusRequest)
}

func (m *LayoutVw) RegisterView(app *config.AppConfig) *LayoutVw{
	log.Println("Registering AppLayoutVw...")
	AppLayoutVw.App = app
	return AppLayoutVw
}

func (m *LayoutVw) RegisterHandler() constants.ViewHandler {
	return &LayoutVw{}
}

// func (m *LayoutVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("[lyoutvw] - Processing request")
// 	CreateLayoutVwData().ProcessHttpRequest(w, r)

// 	// render.RenderModal(w, nil, nil)
// }

// func (m *LayoutVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) any{
// 	CreateLayoutVwData().ProcessMBusRequest(w, r)
// 	return nil
// }

func (m *LayoutVw) HandleRequest(w http.ResponseWriter, event constants.AppEvent) any {
	fmt.Println("[LayoutVw] - HandleRequest")
	var obj LayoutVwData

	if session.SessionSvc.SessionMgr.Exists(event.Context, "layoutvw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Context, "layoutvw").(LayoutVwData)
	} else {
		obj = *CreateLayoutVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Context, "layoutvw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type LayoutVwData struct {
	Base b.BaseTemplateparams
	Data any
	View int
}

type AppLytVwData struct {
	Lbl string
}

func CreateLayoutVwData() *LayoutVwData {
	return &LayoutVwData{
		Base: b.GetBaseTemplateObj(),
		Data: nil,
	}
}

func (m *LayoutVwData) ProcessHttpRequest(w http.ResponseWriter, event constants.AppEvent) *LayoutVwData{
	return m
}

func (m *LayoutVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {}
