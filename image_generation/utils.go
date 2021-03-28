package image_generation

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
	"gopkg.in/ini.v1"
)

func OpenImage(imagepath string) image.Image {
	file, err := ioutil.ReadFile("image_generation/material/" + imagepath)
	if err != nil {
		log.Printf("error for %v", err)
	}
	flagImage, _, decodeErr := image.Decode(bytes.NewReader(file))
	if decodeErr != nil {
		log.Printf("error for %v", err)
	}
	return flagImage
}

func SetFont() *truetype.Font {
	font, err := truetype.Parse(gobold.TTF)
	if err != nil {
		log.Printf("error for %v", err)
	}
	return font
}

func SetSize(size float64, dc *gg.Context) {
	font := SetFont()
	face := truetype.NewFace(font, &truetype.Options{Size: size})
	dc.SetFontFace(face)
}

func GetAvatar(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("error for %v", err)
	}
	defer response.Body.Close()

	file, err := os.Create("image_generation/material/avatar.jpg")
	if err != nil {
		log.Printf("error for %v", err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
}

func GetTime() time.Time {
	t := time.Now()
	cfg, err := ini.Load("app.config")
	if err != nil {
		log.Printf("file=twitter/main.go/13 action=loadcondig error=%v", err)
	}
	if cfg.Section("mode").Key("develop").String() == "on" {
		return t

	}
	t = t.Add(time.Duration(9) * time.Hour)
	return t
}
