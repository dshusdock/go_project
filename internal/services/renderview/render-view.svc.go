package renderview

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/render"
	"dshusdock/go_project/internal/services/jwtauthsvc"
	"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/services/session"
	"dshusdock/go_project/internal/views/base"
	headervw "dshusdock/go_project/internal/views/header"
	"fmt"
	"net/http"
	"time"
)

type RenderView struct {
	App 		*config.AppConfig 
	ViewHandlers 	map[string]con.ViewHandler
}


var RenderViewSvc *RenderView

func MapRenderViewSvc(r *RenderView) {
	RenderViewSvc = r
}

type DisplayData struct {
	Base 		*con.BaseTemplateparams
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
		ViewHandlers: make(map[string]con.ViewHandler),
	}
	RenderViewSvc = obj

	messagebus.GetBus().Subscribe("Event:Click", RenderViewSvc.HandleMBusRequest)
	return RenderViewSvc
	
}

var APP_VIEWS = make (map[string][]string)


func (rv *RenderView) ProcessClickEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing Click Event")
	data := r.PostForm
	fmt.Println("Data: ", data)

	event := con.AppEvent{
		Type: data.Get("type"),
		Label: data.Get("label"),
		Context: r.Context(),
		EventId: con.EVENT_CLICK,
		EventStr: data.Get("view_id") + "_" + data.Get("type") + "_" + data.Get("label"),
	}

	fmt.Println("Event: ", event)

	obj := DisplayData{
		Base: con.GetBaseTemplateObj(event.EventStr),
		Tmplt: make(map[string]*any),
	}

	evt := con.ExtractEventStr(event.EventStr)

	for _, v := range con.APP_VIEWS[evt].Views {
		result := rv.ViewHandlers[v].HandleRequest(w, event)
		obj.Tmplt[v] = &result
	}
	render.RenderAppTemplate(w, r, obj, con.APP_VIEWS[evt].Tmplt)
	fmt.Println("Processing Click Event - Done")
}

func (rv *RenderView) ProcessInit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing Init Event")

	token, _ := jwtauthsvc.CreateToken("test")
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires: time.Now().Add(7 * 24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		Secure: true,
		Name: "token",
		Value: token,
	})

	err := session.SessionSvc.SessionMgr.RenewToken(r.Context())
	if err != nil {
		http.Error(w, "Error renewing token", http.StatusInternalServerError)
		fmt.Println("Error renewing token: ", err)
	}
	
	obj := DisplayData{
		Base: con.GetBaseTemplateObj(""),
		Tmplt: make(map[string]*any),
	}

	event := con.AppEvent{
		Context: r.Context(),
		EventId: con.EVENT_STARTUP,
		EventStr: "startup",
	}

	for _, v := range con.APP_VIEWS[event.EventStr].Views {
		result := rv.ViewHandlers[v].HandleRequest(w, event)
		obj.Tmplt[v] = &result
	}
	render.RenderAppTemplate(w, r, obj, con.APP_VIEWS[event.EventStr].Tmplt)
	fmt.Println("Processing Startup Event - Done")
}


func (rv *RenderView) ProcessRequest(w http.ResponseWriter, r *http.Request, view string) {
	fmt.Println("========================[renderview] - ProcessRequest - [", view, "]========================")

	var rslt any
	var _view int

	//rv.ProcessEvent(w, r, view)

	ev := con.AppEvent{
		Context: r.Context(),
		EventId: con.EVENT_STARTUP,
		
	}
	
	obj := DisplayData{
		Base: con.GetBaseTemplateObj(""),
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
		obj.Base = con.GetBaseTemplateObj("")
	default:
	}	

	if _view == constants.RM_NONE { 
		fmt.Println("No view to render")
		return
	}

	// Now go build the view and render it
	render.RenderAppTemplate(w, r, obj, _view)
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

	eventStr := con.ExtractEventStr(view)
	fmt.Println("EventStr: ", eventStr)

	switch eventStr {
		case "headervw_button_add-item":		
			fmt.Println("Add item button clicked")
			

		case "headervw":
			
		default:
		}

}

