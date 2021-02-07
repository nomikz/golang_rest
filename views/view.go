package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutsDir  = "views/layouts/"
	TemplateDir = "views/"
	TemplateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	formatDirPath(files)

	files = append(
		files,
		layoutFiles()...,
	)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := v.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutsDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func formatDirPath(files []string) {
	for i, dir := range files {
		files[i] = TemplateDir + dir + TemplateExt
	}
}