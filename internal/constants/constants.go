package constants

import (
	"context"
	"net/http"
	
)

const (
	FILESA = iota
	FILESB
	FILESC
)

const (
	EVENT_STARTUP = "Event_Startup"
	EVENT_CLICK  = "Event_Click"
)

const (
	VW_INDEX = iota
)

type EventData struct {
	Id        string
	EventType string
	Event     string
}

type HtmxInfo struct {
	Url string
}

type SubElement struct {
	Type string
	Lbl  string
}

type ViewInterface interface {
	// HandleHttpRequest(w http.ResponseWriter, d url.Values /*ViewInfo*/)
	HandleHttpRequest(w http.ResponseWriter, r *http.Request)
}

type ViewHandler interface {
	HandleRequest(w http.ResponseWriter, event AppEvent) any
	//HandleMBusRequest(w http.ResponseWriter, r *http.Request) any
}

type ViewInfo struct {
	Event   int
	Type    string
	Label   string
	ViewId  string
	ViewStr string
}

type RowData struct {
	Data []string
}

// /////////////Rendered File Map///////////////
const (
	RM_HOME = iota
	RM_NONE
)

type RenderedFileMap struct {
	HOME           []string
	NONE           []string
}

type AppEvent struct {
	ViewId 		string
	Type   		string
	Label  		string
	EventId 	string
	EventStr 	string
	Context 	context.Context

}

func RENDERED_FILE_MAP() *RenderedFileMap {
	return &RenderedFileMap{
		HOME: []string{
			"./ui/html/views/base.tmpl",
			"./ui/html/views/layout.tmpl",
			"./ui/html/views/header.tmpl",
		},
		NONE: []string{
			"",	
		},
	}
}




