package headervw

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/services/session"	
	b "dshusdock/go_project/internal/views/base"
	
	"encoding/gob"

	// "dshusdock/tw_prac1/internal/services/messagebus"
	"fmt"
	"log"
	"net/http"
	// "net/url"
)

type HeaderVw struct {
	App *config.AppConfig
}

var AppHeaderVw *HeaderVw

func init() {
	AppHeaderVw = &HeaderVw{
		App: nil,
	}
	gob.Register(HeaderVwData{})
	messagebus.GetBus().Subscribe("Event:ViewChange", AppHeaderVw.HandleMBusRequest)
}

func (m *HeaderVw) RegisterView(app *config.AppConfig) *HeaderVw{
	log.Println("Registering AppHeaderVw...")
	AppHeaderVw.App = app
	return AppHeaderVw
}

func (m *HeaderVw) RegisterHandler() constants.ViewHandler {
	return &HeaderVw{}
}

func (m *HeaderVw) HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[lyoutvw] - Processing request")
	CreateHeaderVwData().ProcessHttpRequest(w, r)

	// render.RenderModal(w, nil, nil)
}

func (m *HeaderVw) HandleMBusRequest(w http.ResponseWriter, r *http.Request) any{
	CreateHeaderVwData().ProcessMBusRequest(w, r)
	return nil
}

func (m *HeaderVw) HandleRequest(w http.ResponseWriter, r *http.Request) any {
	fmt.Println("[HeaderVw] - HandleRequest")
	var obj HeaderVwData

	if session.SessionSvc.SessionMgr.Exists(r.Context(), "layoutvw") {
		obj = session.SessionSvc.SessionMgr.Pop(r.Context(), "headervw").(HeaderVwData)
	} else {
		obj = *CreateHeaderVwData()	
	}

	obj.ProcessHttpRequest(w, r)	
	session.SessionSvc.SessionMgr.Put(r.Context(), "headervw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type HeaderVwData struct {
	Base b.BaseTemplateparams
	Data any
	View int
}

type AppLytVwData struct {
	Lbl string
}

func CreateHeaderVwData() *HeaderVwData {
	return &HeaderVwData{
		Base: b.GetBaseTemplateObj(),
		Data: nil,
	}
}

func (m *HeaderVwData) ProcessHttpRequest(w http.ResponseWriter, r *http.Request) *HeaderVwData{
	return m
}

func (m *HeaderVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {}
