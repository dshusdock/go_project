package constants

import ()

// /////////////Rendered File Map///////////////
const (
	RM_HOME = iota
	RM_NONE
)

type RenderedFileMap struct {
	HOME           []string
	NONE           []string
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

func GetRenderInfo(viewType int) RenderInfo {
	files := RENDERED_FILE_MAP()

	ri := RenderInfo{
		TemplateName: "",
		TemplateFiles: []string{},
	}

	switch viewType {
	case RM_HOME:
		ri.TemplateName = "base"
		ri.TemplateFiles = files.HOME
	case RM_NONE:
		ri.TemplateName = "none"
		ri.TemplateFiles = files.NONE
	}
	return ri
}
