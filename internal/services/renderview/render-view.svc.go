package renderview

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/render"
	"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/views/base"
	headervw "dshusdock/go_project/internal/views/header"
	"fmt"
	"net/http"
	"net/url"
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
	RenderViewSvc.ViewHandlers["headervw"] = headervw.AppHeaderVw.RegisterHandler()
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

var APP_VIEWS = make (map[string][]string)


func (rv *RenderView) ProcessClickEvent(w http.ResponseWriter, event constants.AppEvent) {
	fmt.Println("Processing Click Event")
	APP_VIEWS["headervw_button_add-item"] = []string{"headervw", "basevw"}

	for _, v := range APP_VIEWS[event.EventStr] {
		_ = rv.ViewHandlers[v].HandleRequest(w, event)
	}
	fmt.Println("Processing Click Event - Done")
}


func (rv *RenderView) ProcessRequest(w http.ResponseWriter, r *http.Request, view string) {
	fmt.Println("========================[renderview] - ProcessRequest - [", view, "]========================")

	var rslt any
	var _view int

	//rv.ProcessEvent(w, r, view)

	ev := constants.AppEvent{
		Context: r.Context(),
		EventId: constants.EVENT_STARTUP,
		
	}
	
	obj := DisplayData{
		Base: base.BaseTemplateparams{},
		Tmplt: make(map[string]*any),
	}

	if (false) { // some special condition) 
		// do something special that returns rslt
	} else {
		rslt = rv.ViewHandlers[view].HandleRequest(w, ev)
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

func (rv *RenderView) ProcessEvent(w http.ResponseWriter, r *http.Request, view string) {
	fmt.Println("========================[renderview] - ProcessEvent - [", view, "]========================")

	d := r.PostForm
	s := d.Get("label")
	fmt.Println("Label: ", s)

	eventStr := createEventStr(view, d)
	fmt.Println("EventStr: ", eventStr)

	switch eventStr {
		case "headervw_button_add-item":		
			fmt.Println("Add item button clicked")
			

		case "headervw":
			
		default:
		}

}

func createEventStr(view string, d url.Values) string {
	return view + "_" + d.Get("type") + "_" + d.Get("label")
}



func (rv *RenderView) RenderTemplate(w http.ResponseWriter, data any, view int) {
	render.RenderTemplate_new(w, nil, data, view)
}

func (rv *RenderView) RegisterHandler(w http.ResponseWriter, data any, view int) {
	
}

