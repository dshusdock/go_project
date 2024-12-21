package constants

import "strings"

type BaseTemplateparams struct {
	LoggedIn 					bool
	DisplayLogin  				bool
	DisplayCreateAccount 		bool
	DisplayCreatAcctResponse 	bool
	SideNav	      				bool
	MainTable	  				bool
	Cards		  				bool
}

func GetBaseTemplateObj(evt string) *BaseTemplateparams{

	switch evt {
	case "headervw_button_add-item":
		return &BaseTemplateparams{
			LoggedIn: true,
			DisplayLogin: false,
			DisplayCreateAccount: false,
			DisplayCreatAcctResponse: false,
			SideNav: true,
			MainTable: true,
			Cards: true,
		}
	default:	
		return &BaseTemplateparams{
		LoggedIn: true,
		DisplayLogin: true,
		DisplayCreateAccount: false,
		DisplayCreatAcctResponse: false,
		SideNav: false,
		MainTable: false,
		Cards: false,
		}
	}
}

var APP_VIEWS = make (map[string]ViewDef)

type ViewDef struct {
	Views 		[]string
	BaseVals 	*BaseTemplateparams
	Tmplt 		int
}

func init() {
	APP_VIEWS["startup"] = ViewDef{Views: []string{"headervw", "basevw"}, BaseVals: nil, Tmplt: RM_HOME}
	APP_VIEWS["headervw_button_add-item"] = ViewDef{Views: []string{"headervw", "basevw"}, BaseVals: nil, Tmplt: RM_HOME}
	APP_VIEWS["headervw_button_login"] = ViewDef{Views: []string{"headervw", "basevw"}, BaseVals: nil, Tmplt: RM_HOME}
}

func ExtractEventStr(str string) string {
	var subKey = ""

	for key := range APP_VIEWS {
		if strings.Compare(str, key) == 0 {
			return key
		}

		if strings.Contains(str, key) {
			subKey = key	
		}
	}
	return subKey
}

