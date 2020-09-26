package match

import (
	"encoding/json"
	"fmt"
	"log"

	"lol-golang/utils"

	"github.com/valyala/fastjson"
)

// Info adds match stats to given channel
func Info(champion int, ids chan int64, results chan PlayerStats) {
	for id := range ids {
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
				//fmt.Println("writing to channel", id)
				results <- s
			}
		}
	}
}
