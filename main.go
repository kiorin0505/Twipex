package main

import (
	"Twipex_project/database"
	"Twipex_project/handler"
	"errors"
	"html/template"
	"io"
	"log"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/ini.v1"
)

type Template struct{}

var templates map[string]*template.Template

var NotFoundError = errors.New("NotFound")

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return templates[name].ExecuteTemplate(w, "base.html", data)
}

func init() {
	loadTemplates()
}

func main() {
	database.DbInit()
	SetSchedule()

	e := echo.New()

	t := &Template{}
	e.Renderer = t

	e.Static("/static", "static")

	store := session.NewCookieStore([]byte("secret-key"))

	handler.SetRoute(e)

	store.MaxAge(86400)
	e.Use(session.Sessions("twipex_session", store))
	e.Use(middleware.Recover())

	cfg, err := ini.Load("app.config")
	if err != nil {
		log.Printf("file=twitter/main.go/41 action=loadconfig error=%v", err)
	}
	e.Logger.Fatal(e.Start(":" + cfg.Section("port").Key("port").String()))

}
