package image_generation

import (
	"Twipex_project/database"
	"Twipex_project/twitter"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type RKDW struct {
	rankname     string
	rp           string
	rpup         string
	kills        string
	killup       string
	damage       string
	damageup     string
	wins         string
	winup        string
	rpupsend     string
	killupsend   string
	damageupsend string
	winupsend    string
	date         string
}

type UserInfo struct {
	id       string
	avatar   string
	rank     string
	platform string
	legend   string
}

func PostImage(now string) {
	log.Printf(now + "work start")
	senddate := GetTime().Format("02")
	queue := database.DbGetPostUser(now)
	for i := range queue {
		db := database.DbGetOne(queue[i].AccountId)
		if senddate != db.Lastsenddate {
			userdata := GetApexData(db.Platform, db.UserId)
			if userdata == nil {
				log.Printf("Failed to get data | platform:%v id:%v twitter:%v", db.Platform, db.UserId, db.AccountName)
			}
			if userdata != nil {
				userdata := userdata[0]
				rkdw := MakeImage(userdata, db)
				if db.Rp != 0 {
					database.DbLogInsert(db.AccountId, rkdw.rp, rkdw.rpup, rkdw.killup, rkdw.damageup, rkdw.winup, GetTime())
				}
				database.DbUpdateUserData(db.AccountId, db.Legend, senddate, rkdw.rankname, rkdw.rp, rkdw.kills, rkdw.damage, rkdw.wins)
				if queue[i].SendInterval == "week" {
					weekday := GetTime().Weekday()
					if weekday == 1 {
						err := twitter.PostTweet(db)
						if err == nil {
							log.Printf("send successfully @%v", db.AccountName)
							database.DbUpdateLastMade(db.AccountId, rkdw.rp, rkdw.kills, rkdw.damage, rkdw.wins, GetTime().Format("2006/01/02"))
						}
					}
					if db.Lastsenddate == "" {
						database.DbUpdateLastMade(db.AccountId, rkdw.rp, rkdw.kills, rkdw.damage, rkdw.wins, GetTime().Format("2006/01/02"))
					}
				} else {
					err := twitter.PostTweet(db)
					if err == nil {
						log.Printf("send successfully @%v", db.AccountName)
						database.DbUpdateLastMade(db.AccountId, rkdw.rp, rkdw.kills, rkdw.damage, rkdw.wins, GetTime().Format("2006/01/02"))
					}
				}
				time.Sleep(time.Second * 2)
			}
		}
	}
	log.Printf(now + "work end")
}

func MakeImage(newdata ApexLawData, beforedata database.UserData) RKDW {
	MakeQrcode(beforedata.AccountId)

	dc := gg.NewContext(1080, 607)
	dc.DrawRectangle(0, 0, 1080, 607)

	flagImage := OpenImage("background/" + beforedata.Legend + ".png")
	dc.DrawImage(flagImage, 0, 0)

	userinfo := UserInfo{
		id:       beforedata.UserId,
		avatar:   newdata.Data.PlatformInfo.AvatarURL,
		rank:     newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
		platform: beforedata.Platform,
		legend:   beforedata.Legend,
	}

	if beforedata.Predator == "on" && newdata.Data.Segments[0].Stats.Rankscore.Value >= 10000 {
		userinfo.rank = "Apex Predator"
	}

	rkdw := CalculateValues(newdata, beforedata)

	area := OpenImage("area.png")
	dc.DrawImage(area, 0, 0)
	UserInfoGenerator(userinfo, dc)

	DrawRp(rkdw, dc)
	DrawValues("Kills", rkdw.kills, rkdw.killup, dc)
	DrawValues("Damege", rkdw.damage, rkdw.damageup, dc)
	if beforedata.Winad != "" {
		DrawValues("Wins", rkdw.wins, rkdw.winup, dc)
	}

	DrawDate(dc, beforedata)

	dc.SavePNG("data.png")

	return rkdw
}

func DrawDate(dc *gg.Context, beforedata database.UserData) {
	dc.SetRGB(0.85, 0.85, 0.85)
	SetSize(24, dc)
	today := GetTime().Format("2006/01/02")
	yesterday := GetTime().AddDate(0, 0, -1).Format("2006/01/02")

	if beforedata.SendInterval == "week" {
		if beforedata.LastMadeDate == yesterday || beforedata.LastMadeDate == "" {
			dc.DrawString(today, 920, 595)
		} else {
			dc.DrawString(beforedata.LastMadeDate+" ~ "+today, 800, 595)
		}
	} else {
		dc.DrawString(today, 920, 595)
	}
}

func MakeQrcode(twitterid string) {
	qrCode, _ := qr.Encode("https://twipex.herokuapp.com//data/"+twitterid, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("image_generation/material/qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}

func UserInfoGenerator(info UserInfo, dc *gg.Context) {
	GetAvatar(info.avatar)
	avatar := OpenImage("avatar.jpg")
	if avatar == nil {
		avatar = OpenImage("avatar_failed.jpg")
	}
	platform := OpenImage(info.platform + "48.png")
	rank := OpenImage("rank/" + info.rank + ".jpg")
	dc.DrawImage(resize.Resize(156, 156, rank, resize.Lanczos3), 274, 44)
	dc.DrawImage(resize.Resize(208, 208, avatar, resize.Lanczos3), 20, 20)
	dc.DrawImage(platform, 476, 60)

	SetFont()
	SetSize(20, dc)
	dc.SetRGB(0.85, 0.85, 0.85)
	if info.platform == "origin" {
		dc.DrawString("Origin", 524, 100)
	} else if info.platform == "psn" {
		dc.DrawString("PSN", 524, 100)
	} else if info.platform == "xbl" {
		dc.DrawString("XBOX", 524, 100)
	}
	SetSize(60, dc)
	dc.DrawString(info.id, 476, 165)

	legendbunner := OpenImage("tiles/" + info.legend + ".png")
	dc.DrawImage(resize.Resize(250, 279, legendbunner, resize.Lanczos3), 240, 275)
	flagImage := OpenImage("qrcode.png")
	dc.SetRGB(0.85, 0.85, 0.85)
	SetSize(38, dc)
	dc.DrawString(info.legend, 40, 320)
	dc.SetRGB(0.75, 0.75, 0.75)
	SetSize(20, dc)
	dc.DrawString("Active Legend", 40, 280)
	dc.DrawString("↑ Check details", 52, 520)

	dc.DrawImage(resize.Resize(140, 140, flagImage, resize.Lanczos3), 54, 350)
}

func DrawRp(r RKDW, dc *gg.Context) {
	dc.SetRGB(0.85, 0.85, 0.85)
	SetSize(50, dc)

	dc.DrawString(r.rp, 360, 280)

	dc.SetRGB(0.75, 0.75, 0.75)
	dc.DrawString("RP", 280, 280)
	SetSize(20, dc)
	dc.DrawString("To the next rank", 700, 280)
	SetSize(30, dc)
	nextrank := strconv.Itoa(CalculateNextRank(r))
	if nextrank == "0" {
		nextrank = "-"
	}
	dc.DrawString(nextrank, 875, 280)

	if r.rpup != "0" && r.rpupsend != r.rp {
		SetSize(40, dc)
		val, _ := strconv.Atoi(r.rpup)
		if val > 0 {
			dc.SetRGB(0.8, 0.3, 0.3)
			dc.DrawString(r.rpupsend, 545, 280)
			flagImage := OpenImage("up32.png")
			dc.DrawImage(flagImage, 510, 250)
		} else {
			dc.SetRGB(0.25, 0.67, 0.83)
			dc.DrawString(r.rpupsend, 545, 280)
			flagImage := OpenImage("down32.png")
			dc.DrawImage(flagImage, 510, 250)
		}

	}
}

func DrawValues(valuename, value, rise string, dc *gg.Context) {
	dc.SetRGB(0.75, 0.75, 0.75)
	SetSize(40, dc)
	dc.DrawString(valuename, 490, 440)
	dc.SetRGB(0.85, 0.85, 0.85)
	dc.DrawString(value, 660, 440)

	if rise != "0" && rise != value {
		SetSize(40, dc)
		dc.SetRGB(0.8, 0.3, 0.3)
		dc.DrawString(rise, 885, 440)
		flagImage := OpenImage("up32.png")
		dc.DrawImage(flagImage, 850, 410)
	}
}
