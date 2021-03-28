package handler

import (
	"Twipex_project/database"
	"Twipex_project/twitter"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

type IndexData struct {
	LegendName  string
	Login       bool
	TwitterName string
}

type LogPackage struct {
	LogData     []database.PlayerLogData
	LogUser     string
	AccountName string
	AccountId   string
}

func SetRoute(e *echo.Echo) {
	e.GET("/", Index)
	e.GET("/howto", HowtoUse)
	e.GET("/logout", Logout)
	e.GET("/oauth", twitter.Oauth)
	e.GET("/callback", twitter.Callback)
	e.GET("/auth/twitter/callback", twitter.Callback)
	e.GET("/confirm", Confirm)
	e.POST("/create", UpdateProfile)
	e.GET("/setting", Setting)
	e.GET("/data/:id", UserData)
	e.GET("/contact", Contact)
	e.POST("/postmessage", PostMessage)
}

func LoginCertficate(c echo.Context) bool {
	session := session.Default(c)
	twitterid := session.Get("twitter_id")
	if twitterid == nil {
		return false
	}
	return true
}

func LoginData(c echo.Context) database.UserData {
	session := session.Default(c)
	id := session.Get("twitter_id").(string)
	return database.DbGetOne(id)
}

func Index(c echo.Context) error {
	var logindata database.UserData
	if LoginCertficate(c) == true {
		logindata = LoginData(c)
		return c.Render(http.StatusOK, "index", logindata)
	}
	return c.Render(http.StatusOK, "index", logindata)
}

func Setting(c echo.Context) error {
	if LoginCertficate(c) == false {
		return c.Redirect(http.StatusFound, "/")
	}
	logindata := LoginData(c)
	return c.Render(http.StatusOK, "setting", logindata)
}

func Contact(c echo.Context) error {
	var logindata database.UserData
	if LoginCertficate(c) == true {
		logindata = LoginData(c)
		return c.Render(http.StatusOK, "contact", logindata)
	}
	return c.Render(http.StatusOK, "contact", logindata)
}

func HowtoUse(c echo.Context) error {
	var logindata database.UserData
	if LoginCertficate(c) == true {
		logindata = LoginData(c)
		return c.Render(http.StatusOK, "howtouse", logindata)
	}
	return c.Render(http.StatusOK, "howtouse", logindata)
}

func Confirm(c echo.Context) error {
	var logindata database.UserData
	if LoginCertficate(c) == true {
		logindata = LoginData(c)
		return c.Render(http.StatusOK, "confirm.html", logindata)
	}
	return c.Render(http.StatusOK, "confirm.html", logindata)
}

func UserData(c echo.Context) error {
	var logpackage LogPackage
	id := c.Param("id")
	logdata := database.DbLogGet(id)
	loguser := database.DbGetOne(id).UserId
	logpackage.LogUser = loguser
	logpackage.LogData = logdata
	if LoginCertficate(c) == true {
		logpackage.AccountName = LoginData(c).AccountName
		logpackage.AccountId = LoginData(c).AccountId
		if len(logdata) == 0 {
			return c.Render(http.StatusFound, "nodata", logpackage)
		}
		return c.Render(http.StatusOK, "data", logpackage)
	}
	if len(logdata) == 0 {
		return c.Render(http.StatusFound, "nodata.html", logpackage)
	}
	return c.Render(http.StatusFound, "data", logpackage)
}

func UpdateProfile(c echo.Context) error {
	session := session.Default(c)
	twitterid := session.Get("twitter_id").(string)
	platform := c.FormValue("platform")
	id := c.FormValue("id")
	legend := c.FormValue("legend")
	winad := c.FormValue("winad")
	sendtime := c.FormValue("time")
	sendinterval := c.FormValue("sendinterval")
	predator := c.FormValue("predator")
	database.DbUpdateProfile(twitterid, platform, id, legend, winad, sendtime, sendinterval, predator)
	return c.Redirect(http.StatusFound, "/")
}

func Logout(c echo.Context) error {
	session := session.Default(c)
	session.Delete("twitter_id")
	session.Save()
	return c.Redirect(http.StatusFound, "/")
}

func PostMessage(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	content := c.FormValue("content")
	database.DbCreateMessage(name, address, content)
	return c.Redirect(http.StatusFound, "/")
}
