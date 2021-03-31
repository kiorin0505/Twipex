package twitter

import (
	"Twipex_project/database"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"gopkg.in/ini.v1"
)

func getConnect() *oauth.Client {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Printf("file=twitter/main.go/13 action=loadcondig error=%v", err)
	}
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  cfg.Section("twitter").Key("consumer_token").String(),
			Secret: cfg.Section("twitter").Key("consumer_secret").String(),
		},
	}
}

func getAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	oc := getConnect()
	at, _, err := oc.RequestToken(nil, rt, oauthVerifier)

	return at, err
}

type Data struct {
	id   string
	name string
}

func getSelfData(token, secret string) Data {

	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Printf("file=twitter/main.go/41 action=loadconfig error=%v", err)
	}
	anaconda.SetConsumerKey(cfg.Section("twitter").Key("consumer_token").String())
	anaconda.SetConsumerSecret(cfg.Section("twitter").Key("consumer_secret").String())
	api := anaconda.NewTwitterApi(token, secret)
	data, _ := api.GetSelf(nil)
	twitterdata := Data{
		id:   data.IdStr,
		name: data.ScreenName,
	}

	return twitterdata
}

func Oauth(c echo.Context) error {
	config := getConnect()
	rt, err := config.RequestTemporaryCredentials(nil, "http://localhost:8080/callback", nil)
	if err != nil {
		log.Printf("file=twitter/oauth.go/13 action=requestcredential error=%v", err)
	}
	session := session.Default(c)
	session.Set("request_token", rt.Token)
	session.Set("request_token_secret", rt.Secret)
	session.Save()

	url := config.AuthorizationURL(rt, nil)

	return c.Redirect(http.StatusFound, url)
}

func Callback(c echo.Context) error {
	session := session.Default(c)
	secret := c.QueryParam("oauth_verifier")
	at, err := getAccessToken(
		&oauth.Credentials{
			Token:  session.Get("request_token").(string),
			Secret: session.Get("request_token_secret").(string),
		},
		secret,
	)
	if err != nil {
		//log.Printf("file=callback.go/24 action=getAccessToken error=%v", err)
		return c.Redirect(http.StatusFound, "/")
	}

	twidata := getSelfData(at.Token, at.Secret)

	session.Set("twitter_id", twidata.id)
	session.Set("twitter_name", twidata.name)

	session.Save()
	if database.Check(twidata.id) {
		return c.Redirect(http.StatusFound, "/")
	}
	database.InitInsert(at.Token, at.Secret, twidata.id, twidata.name)
	return c.Redirect(http.StatusFound, "/setting")
}

func PostTweet(db database.UserData) error {

	content := db.UserId + "'s " + db.Legend + " Status made by #Twipex  For more details https://twipex.herokuapp.com/data/" + db.AccountId

	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Printf("file=twitter/post.go/19 action=loadconfig error=%v", err)
	}
	anaconda.SetConsumerKey(cfg.Section("twitter").Key("consumer_token").String())
	anaconda.SetConsumerSecret(cfg.Section("twitter").Key("consumer_secret").String())
	api := anaconda.NewTwitterApi(db.Token, db.Secret)

	file, _ := os.Open("data.png")
	fileData, _ := ioutil.ReadAll(file)

	base64String := base64.StdEncoding.EncodeToString(fileData)

	media, _ := api.UploadMedia(base64String)

	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)

	_, err = api.PostTweet(content, v)
	if err != nil {
		log.Printf("file=twitter/post.go/37 action=posttweet error=%v", err)
	}
	return err
}
