package handler

import (
	"Twipex_project/database"
	"Twipex_project/twitter"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

type logPackage struct {
	LogData     []database.PlayerLogData
	LogUser     string
	AccountName string
	AccountId   string
}

func SetRoute(e *echo.Echo) {
	e.GET("/", index)
	e.GET("/howto", howtoUse)
	e.GET("/logout", logout)
	e.GET("/oauth", twitter.Oauth)
	e.GET("/callback", twitter.Callback)
	e.GET("/auth/twitter/callback", twitter.Callback)
	e.GET("/confirm", confirm)
	e.POST("/create", updateProfile)
	e.GET("/setting", setting)
	e.GET("/data/:id", userData)
	e.GET("/contact", contact)
	e.POST("/postmessage", postMessage)
}

func loginCertficate(c echo.Context) bool {
	session := session.Default(c)
	twitterid := session.Get("twitter_id")
	if twitterid == nil {
		return false
	}
	return true
}

func loginData(c echo.Context) database.UserData {
	session := session.Default(c)
	id := session.Get("twitter_id").(string)
	return database.GetOne(id)
}

func index(c echo.Context) error {
	var logindata database.UserData
	if loginCertficate(c) == true {
		logindata = loginData(c)
		return c.Render(http.StatusOK, "index", logindata)
	}
	return c.Render(http.StatusOK, "index", logindata)
}

func setting(c echo.Context) error {
	if loginCertficate(c) == false {
		return c.Redirect(http.StatusFound, "/")
	}
	logindata := loginData(c)
	return c.Render(http.StatusOK, "setting", logindata)
}

func contact(c echo.Context) error {
	var logindata database.UserData
	if loginCertficate(c) == true {
		logindata = loginData(c)
		return c.Render(http.StatusOK, "contact", logindata)
	}
	return c.Render(http.StatusOK, "contact", logindata)
}

func howtoUse(c echo.Context) error {
	var logindata database.UserData
	if loginCertficate(c) == true {
		logindata = loginData(c)
		return c.Render(http.StatusOK, "howtouse", logindata)
	}
	return c.Render(http.StatusOK, "howtouse", logindata)
}

func confirm(c echo.Context) error {
	var logindata database.UserData
	if loginCertficate(c) == true {
		logindata = loginData(c)
		return c.Render(http.StatusOK, "confirm.html", logindata)
	}
	return c.Render(http.StatusOK, "confirm.html", logindata)
}

func userData(c echo.Context) error {
	var logpackage logPackage
	id := c.Param("id")
	logdata := database.LogGet(id)
	loguser := database.GetOne(id).UserId
	logpackage.LogUser = loguser
	logpackage.LogData = logdata
	if loginCertficate(c) == true {
		logpackage.AccountName = loginData(c).AccountName
		logpackage.AccountId = loginData(c).AccountId
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

func updateProfile(c echo.Context) error {
	session := session.Default(c)
	twitterid := session.Get("twitter_id").(string)
	platform := c.FormValue("platform")
	id := c.FormValue("id")
	legend := c.FormValue("legend")
	winad := c.FormValue("winad")
	sendtime := c.FormValue("time")
	sendinterval := c.FormValue("sendinterval")
	predator := c.FormValue("predator")
	database.UpdateProfile(twitterid, platform, id, legend, winad, sendtime, sendinterval, predator)
	return c.Redirect(http.StatusFound, "/")
}

func logout(c echo.Context) error {
	session := session.Default(c)
	session.Delete("twitter_id")
	session.Save()
	return c.Redirect(http.StatusFound, "/")
}

func postMessage(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	content := c.FormValue("content")
	database.CreateMessage(name, address, content)
	return c.Redirect(http.StatusFound, "/")
}
