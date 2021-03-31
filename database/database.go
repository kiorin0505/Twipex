package database

import (
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

type UserData struct {
	gorm.Model
	//ApexData
	UserId   string
	Platform string
	Rp       int
	Kills    int
	Damage   int
	Wins     int
	Rank     string

	//TwitterData
	Token       string
	Secret      string
	AccountId   string
	AccountName string

	//TwipexData
	Winad        string
	SendTime     string
	SendInterval string
	Lastsenddate string
	Legend       string
	BeforeLegend string
	Predator     string

	LastMadeRp     int
	LastMadeKills  int
	LastMadeDamage int
	LastMadeWins   int
	LastMadeDate   string
}

type PlayerLogData struct {
	gorm.Model
	AccountId string
	Rp        int
	RpUp      int
	KillUp    int
	WinsUp    int
	DamageUp  int
	Year      int
	Month     int
	Date      int
}

type contactData struct {
	gorm.Model
	Name    string
	Address string
	Content string
}

func Open() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./database/test.sqlite3")
	if err != nil {
		log.Panicf("Failed to Open DataBase: %v", err)
	}
	return db
}

func Init() {
	db := Open()
	db.AutoMigrate(&UserData{})
	db.AutoMigrate(&PlayerLogData{})
	db.AutoMigrate(&contactData{})
	defer db.Close()
}

func GetPostUser(now string) []UserData {
	db := Open()
	var userdata []UserData
	db.Find(&userdata, "send_time = ?", now)
	db.Close()
	return userdata
}

func GetOne(twitterid string) UserData {
	db := Open()
	var userdata UserData
	db.First(&userdata, "account_id = ?", twitterid)
	db.Close()
	return userdata
}

func InitInsert(token, secret, twitterid, twittername string) {
	db := Open()
	db.Create(&UserData{AccountId: twitterid, AccountName: twittername, Token: token, Secret: secret})
	defer db.Close()

}

func UpdateProfile(twitterid, platform, id, legend, winad, time, sendinterval, predator string) {
	db := Open()
	var userdata UserData
	db.First(&userdata, "account_id = ?", twitterid)
	userdata.Platform = platform
	userdata.UserId = id
	userdata.Legend = legend
	userdata.Winad = winad
	userdata.SendTime = time
	userdata.SendInterval = sendinterval
	userdata.Predator = predator
	db.Save(&userdata)
	db.Close()
}

func UpdateUserData(twitterid, beforelegend, lastsenddate, rankname, rp, kills, damage, wins string) {
	db := Open()
	var userdata UserData
	db.First(&userdata, "account_id = ?", twitterid)
	userdata.Rank = rankname
	userdata.Rp, _ = strconv.Atoi(rp)
	userdata.Kills, _ = strconv.Atoi(kills)
	userdata.Damage, _ = strconv.Atoi(damage)
	userdata.Wins, _ = strconv.Atoi(wins)
	userdata.BeforeLegend = beforelegend
	userdata.Lastsenddate = lastsenddate
	db.Save(&userdata)
	defer db.Close()
}

func Check(twitterid string) bool {
	db := Open()
	var userdata UserData
	db.First(&userdata, "account_id = ?", twitterid)
	if userdata.AccountId == "" {
		return false
	}
	return true
}

func LogInsert(twitterid, rp, rpup, killup, damageup, winsup string, makedate time.Time) {
	db := Open()
	t := makedate
	intrp, _ := strconv.Atoi(rp)
	intrpup, _ := strconv.Atoi(rpup)
	intkillup, _ := strconv.Atoi(killup)
	intdamageup, _ := strconv.Atoi(damageup)
	intwinsup, _ := strconv.Atoi(winsup)
	year, _ := strconv.Atoi(t.Format("2006"))
	month, _ := strconv.Atoi(t.Format("01"))
	month = month - 1
	date, _ := strconv.Atoi(t.Format("02"))
	db.Create(&PlayerLogData{AccountId: twitterid, Rp: intrp, RpUp: intrpup, KillUp: intkillup, DamageUp: intdamageup, WinsUp: intwinsup, Year: year, Month: month, Date: date})
	defer db.Close()
}

func LogGet(twitterid string) []PlayerLogData {
	db := Open()
	var playerlogdata []PlayerLogData
	db.Find(&playerlogdata, "account_id = ?", twitterid)
	defer db.Close()
	return playerlogdata
}

func CreateMessage(name, address, content string) {
	db := Open()
	db.Create(&contactData{Name: name, Address: address, Content: content})
	defer db.Close()
}

func UpdateLastMade(twitterid, lastrp, lastkills, lastdamage, lastwins, lastdate string) {
	db := Open()
	var userdata UserData
	db.First(&userdata, "account_id = ?", twitterid)
	userdata.LastMadeRp, _ = strconv.Atoi(lastrp)
	userdata.LastMadeKills, _ = strconv.Atoi(lastkills)
	userdata.LastMadeDamage, _ = strconv.Atoi(lastdamage)
	userdata.LastMadeWins, _ = strconv.Atoi(lastwins)
	userdata.LastMadeDate = lastdate

	db.Save(&userdata)
	defer db.Close()
}
