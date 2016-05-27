package gotpl

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

type Helper struct {
	templates map[string]*template.Template
}

// New returns a Helper object
// path: path of the template directory
func New(path string) (*Helper, error) {
	var helper Helper = Helper{}
	path = strings.TrimSuffix(path, "/")
	shares, err := filepath.Glob(path + "/shares/*.tpl")
	if err != nil {
		return nil, err
	}
	tpls, err := filepath.Glob(path + "/*.tpl")
	if err != nil {
		return nil, err
	}

	temps := make(map[string]*template.Template)

	for _, tpl := range tpls {
		files := append(shares, tpl)
		temps[filepath.Base(tpl)] = template.Must(template.ParseFiles(files...))
	}

	helper.templates = temps
	return &helper, nil
}

// name: name of the template, for example: "index" (without extension ".tpl")
func (helper *Helper) Render(w http.ResponseWriter, name string, data map[string]interface{}) error {
	tpl, ok := helper.templates[name + ".tpl"]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	return tpl.ExecuteTemplate(w, "layout.tpl", data)
}