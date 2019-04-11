package templating

import "html/template"

var err error
var tpls *template.Template

// InitTemplating function is used to initialize the templating
func InitTemplating(path ...string) error {
	tpls, err = template.ParseGlob(path[0])
	if err != nil {
		return err
	}

	if len(path) > 1 {
		for _, p := range path {
			tpls, err = tpls.ParseGlob(p)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetTpls function is used to get the tpls variable from other packages
func GetTpls() *template.Template {
	return tpls
}
