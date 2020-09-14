package match

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/meximonster/lol-golang/utils"
	"github.com/valyala/fastjson"
)

// Info adds match stats to given channel
func Info(id int64, champion int, resultsChan chan (PlayerStats)) {
	var s PlayerStats
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
			l1 := vv[i].MarshalTo(nil)
			err := json.Unmarshal(l1, &s)
			if err != nil {
				log.Fatal(err)
			}
			resultsChan <- s
		}
	}
}
