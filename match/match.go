package match

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/meximonster/lol-golang/utils"
	"github.com/valyala/fastjson"
)

// Info returns player stats from a match given matchId
func Info(id int64, champion int, resultsChan chan (Stats)) {
	var s Stats
	var t Timeline
	url := "https://euw1.api.riotgames.com/lol/match/v4/matches/" + fmt.Sprint(id)
	body := utils.GetReq(url)
	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	if err != nil {
		log.Fatal(err)
	}
	vv := v.GetArray("participants")
	for i := range vv {
		if vv[i].GetInt("championId") == champion {
			l1 := vv[i].Get("stats").MarshalTo(nil)
			err := json.Unmarshal(l1, &s)
			if err != nil {
				log.Fatal("Could not unmarshal json", err)
			}
			l2 := vv[i].Get("timeline").MarshalTo(nil)
			err = json.Unmarshal(l2, &t)
			if err != nil {
				log.Fatal(err)
			}
			// fmt.Println(s.Kills, s.Deaths, s.Assists)
			resultsChan <- s
		}
	}
}
