package match

import (
	"fmt"
	"log"

	"github.com/meximonster/lol-golang/utils"
	"github.com/valyala/fastjson"
)

// Info returns player stats from a match given matchId
func Info(id int64, champion int) {
	url := "https://euw1.api.riotgames.com/lol/match/v4/matches/" + fmt.Sprint(id)
	body := utils.GetReq(url)
	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	if err != nil {
		log.Fatal("Could not parse body", err)
	}
	vv := v.GetArray("participants")
	for i := range vv {
		if vv[i].GetInt("championId") == champion {
			s := vv[i].Get("stats").MarshalTo(nil)
			fmt.Println(string(s))
		}
	}
}
