package renderview

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/render"
	"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/views/base"
	"fmt"
	"net/http"
)

type RenderView struct {
	App 		*config.AppConfig 
	ViewHandlers 	map[string]constants.ViewHandler
}

var RenderViewSvc *RenderView

func MapRenderViewSvc(r *RenderView) {
	RenderViewSvc = r
}

type DisplayData struct {
	Base 		base.BaseTemplateparams
	Tmplt   	map[string]*any
	TestStr 	string
}

func InitRouteHandlers() {
	// Register the views
	RenderViewSvc.ViewHandlers["basevw"] = base.AppBaseVw.RegisterHandler()
}

func NewRenderViewSvc(app *config.AppConfig) *RenderView {
	
	obj := &RenderView{
		App: app,
		ViewHandlers: make(map[string]constants.ViewHandler),
	}
	RenderViewSvc = obj

	messagebus.GetBus().Subscribe("Event:Click", RenderViewSvc.HandleMBusRequest)
	return RenderViewSvc
	
}

func (rv *RenderView) ProcessRequest(w http.ResponseWriter, r *http.Request, view string) {
	var rslt any
	var _view int
	
	obj := DisplayData{
		Base: base.BaseTemplateparams{},
		Tmplt: make(map[string]*any),
	}

	if (false) { // some special condition) 
		// do something special that returns rslt
	} else {
		rslt = rv.ViewHandlers[view].HandleRequest(w, r)
	}
	obj.Tmplt[view] = &rslt

	switch view {
	case "basevw":		
		_view = rslt.(base.BaseVwData).View
		obj.Base = rslt.(base.BaseVwData).Base
	default:
	}	

	if _view == constants.RM_NONE { 
		fmt.Println("No view to render")
		return
	}

	// Now go build the view and render it
	rv.RenderTemplate(w, obj, _view)
}

func (rv *RenderView) HandleMBusRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[renderview] - HandleMBusRequest")
    d := r.PostForm
	id := d.Get("view_id")	

	switch id {
	case "basevw":
	}
}

func (rv *RenderView) RenderTemplate(w http.ResponseWriter, data any, view int) {
	render.RenderTemplate_new(w, nil, data, view)
}

