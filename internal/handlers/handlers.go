package handlers

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/services/renderview"
	//"dshusdock/go_project/internal/views/base"
	//"dshusdock/go_project/internal/views/layoutvw"
	"fmt"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	MBS messagebus.MessageBusSvc
}

func initRouteHandlers() {
	// Register the views
	//Repo.App.ViewCache["basevw"] = base.AppBaseVw.RegisterView(Repo.App)
	//Repo.App.ViewCache["lyoutvw"] = layoutvw.AppLayoutVw.RegisterView(Repo.App)

}

// http.ResponseWriter, r *http.Request NewRepo creates a new repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
	initRouteHandlers()
}

func (m *Repository) Base(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Base Handler PATH-", r.URL.Path)

	// ptr := m.App.ViewCache["loginvw"]
	// ptr.HandleHttpRequest(w, r)	

	renderview.RenderViewSvc.ProcessInit(w, r)

}

func (m *Repository) HandleClickEvents(w http.ResponseWriter, r *http.Request) {
	val := m.App.SessionManager.Get(r.Context(), "LoggedIn")	
	fmt.Println("[Handlers] Logged In -  ", val)
	
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := r.PostForm
	data.Add("event", constants.EVENT_CLICK)
	v_id := data.Get("view_id")

	if v_id == "" {
		_ = fmt.Errorf("no handler for route")
		return
	}

	renderview.RenderViewSvc.ProcessClickEvent(w, r)
}


