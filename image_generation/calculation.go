package image_generation

import (
	"Twipex_project/database"
	"strconv"
)

func CalculateValues(newdata ApexLawData, beforedata database.UserData) RKDW {
	for i := 0; i < len(newdata.Data.Segments); i++ {
		if newdata.Data.Segments[i].Metadata.Name == beforedata.Legend {
			if beforedata.Legend != beforedata.BeforeLegend {
				rpchange := int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.Rp)
				rpchangesend := int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.LastMadeRp)
				rkdw := RKDW{
					rankname:     newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
					rp:           strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value)),
					rpup:         strconv.Itoa(rpchange),
					kills:        strconv.Itoa(int(newdata.Data.Segments[i].Stats.Kills.Value)),
					killup:       "0",
					damage:       strconv.Itoa(int(newdata.Data.Segments[i].Stats.Damage.Value)),
					damageup:     "0",
					wins:         strconv.Itoa(int(newdata.Data.Segments[i].Stats.Wins.Value)),
					winup:        "0",
					rpupsend:     strconv.Itoa(rpchangesend),
					killupsend:   "0",
					damageupsend: "0",
					winupsend:    "0",
					date:         GetTime().Format("2006/01/02"),
				}
				return rkdw
			}

			rpchange := int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.Rp)
			killchange := int(newdata.Data.Segments[i].Stats.Kills.Value) - (beforedata.Kills)
			damagechange := int(newdata.Data.Segments[i].Stats.Damage.Value) - (beforedata.Damage)
			winchange := int(newdata.Data.Segments[i].Stats.Wins.Value) - (beforedata.Wins)
			rpchangesend := int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.LastMadeRp)
			killchangesend := int(newdata.Data.Segments[i].Stats.Kills.Value) - (beforedata.LastMadeKills)
			damagechangesend := int(newdata.Data.Segments[i].Stats.Damage.Value) - (beforedata.LastMadeDamage)
			winchangesend := int(newdata.Data.Segments[i].Stats.Wins.Value) - (beforedata.LastMadeWins)
			rkdw := RKDW{
				rankname:     newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
				rp:           strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value)),
				rpup:         strconv.Itoa(rpchange),
				kills:        strconv.Itoa(int(newdata.Data.Segments[i].Stats.Kills.Value)),
				killup:       strconv.Itoa(killchange),
				damage:       strconv.Itoa(int(newdata.Data.Segments[i].Stats.Damage.Value)),
				damageup:     strconv.Itoa(damagechange),
				wins:         strconv.Itoa(int(newdata.Data.Segments[i].Stats.Wins.Value)),
				winup:        strconv.Itoa(winchange),
				rpupsend:     strconv.Itoa(rpchangesend),
				killupsend:   strconv.Itoa(killchangesend),
				damageupsend: strconv.Itoa(damagechangesend),
				winupsend:    strconv.Itoa(winchangesend),
				date:         GetTime().Format("2006/01/02"),
			}
			return rkdw
		}
	}

	//選択したレジェンドのデータがなかったときの処理
	rpchange := int(newdata.Data.Segments[0].Stats.Rankscore.Value) - (beforedata.Rp)
	rkdw := RKDW{
		rankname: newdata.Data.Segments[0].Stats.Rankscore.Rankmetadata.Rankname,
		rp:       strconv.Itoa(int(newdata.Data.Segments[0].Stats.Rankscore.Value)),
		rpup:     strconv.Itoa(rpchange),
		kills:    "0",
		killup:   "0",
		damage:   "0",
		damageup: "0",
		wins:     "0",
		winup:    "0",
	}
	return rkdw
}

func CalculateNextRank(r RKDW) int {
	val, _ := strconv.Atoi(r.rp)
	if val < 300 {
		return 300 - val
	}
	if val < 600 {
		return 600 - val
	}
	if val < 900 {
		return 900 - val
	}
	if val < 1200 {
		return 1200 - val
	}
	if val < 1600 {
		return 1600 - val
	}
	if val < 2000 {
		return 2000 - val
	}
	if val < 2400 {
		return 2400 - val
	}
	if val < 2800 {
		return 2800 - val
	}
	if val < 3300 {
		return 3300 - val
	}
	if val < 3800 {
		return 3800 - val
	}
	if val < 4300 {
		return 4300 - val
	}
	if val < 4800 {
		return 4800 - val
	}
	if val < 5400 {
		return 5400 - val
	}
	if val < 6000 {
		return 6000 - val
	}
	if val < 6600 {
		return 6600 - val
	}
	if val < 7200 {
		return 7200 - val
	}
	if val < 7900 {
		return 7900 - val
	}
	if val < 8600 {
		return 8600 - val
	}
	if val < 9300 {
		return 9300 - val
	}
	if val < 10000 {
		return 10000 - val
	}
	if val >= 10000 {
		return 0
	}
	return 0
}
