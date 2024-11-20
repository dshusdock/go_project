package render

import (
	"dshusdock/go_project/internal/constants"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var files = []string{
	"./ui/html/pages/base.tmpl.html",
	"./ui/html/pages/layout.tmpl.html",
	"./ui/html/pages/header.tmpl.html",
	"./ui/html/pages/test/page1.tmpl.html",
	"./ui/html/pages/sidenav.tmpl.html",
	"./ui/html/pages/system-list.tmpl.html",
	"./ui/html/pages/test-modal.tmpl.html",
}

type Payload struct {
    Server string
}

func JSONResponse(w http.ResponseWriter, data string) {
	t := Payload{Server: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, d any) {
	tmpl := template.Must(template.ParseFiles(files...))

	tmpl.ExecuteTemplate(w, "base", d)
}

// RenderTemplate renders a template
func RenderTemplate_new(w http.ResponseWriter, r *http.Request, data any, _type int) {
	ptr := constants.RENDERED_FILE_MAP()
	fmt.Println("In RenderTemplate_new - Type: ", _type)

	var tmplFiles []string
	var tmplName string

	switch _type {
	case constants.RM_HOME:
		tmplFiles = ptr.HOME
		tmplName = "base"
	
	case constants.RM_NONE:
		tmplFiles = ptr.NONE
		tmplName = ""
	default:
		tmplFiles = ptr.HOME
		tmplName = "base"
	}
	template.Must(template.ParseFiles(tmplFiles...)).ExecuteTemplate(w, tmplName, data)
}
