package main

import (
	"html/template"
)

func loadTemplates() {
	var baseTemplate = "templates/base.html"
	templates = make(map[string]*template.Template)
	templates["index"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/index.html"))
	templates["setting"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/setting.html"))
	templates["data"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/dashboard.html"))
	templates["howtouse"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/howtouse.html"))
	templates["contact"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/contact.html"))
	templates["nodata"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/nodata.html"))
}
