package player

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fastjson"
	"lol-golang/utils"
)

type summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

var m = map[string]string{
	"xtiltos":        "VIlzlZHXTczJ_BEM2442GVcE91TZr-yZoa0_Vo540FumNQ",
	"GiannisVrettos": "UoDNZTMsW7CLWaWAITBYBjUwppEpGJ0yLpO2x06aSVyhWg",
	"Treloumination": "aqjxoVOY6q0braazLQS3nI2qELtCMY9tW0rwJl6gWAfFVQ",
	"claymanjf":      "G06KSKh9wVgsX_rQ_01UsfuecgKTui9RtuSim-Eodt0cCA",
	"fullonwokey":    "vNfGjvyV26D-U543RouMBnEdlr46HN90GrqCtO50_icmNw",
	"Dragonlet":      "vwtgv_ljF3OUVCFwchhJ_5e8vDPAbjD-misNvpQHQtHZew",
}

// GetaccID returns the encrypted account id of a summoner
func GetaccID(name string) string {
	v, ok := m[name]
	if ok {
		return v
	}
	v = AccInfo(name)
	return v
}

// AccInfo is a helper function for GetaccID
func AccInfo(name string) string {
	url := "https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/" + name
	body := utils.GetReq(url)
	var s summoner
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal("Could not unmarshal body", err)
	}
	return s.AccountID
}

// GetMatches returns a list with the gameIds of a player for a specific champion
func GetMatches(id string, c string) []int64 {
	var m []int64
	params := "?champion=" + c + "&season=13"
	url := "https://euw1.api.riotgames.com/lol/match/v4/matchlists/by-account/" + id + params
	body := utils.GetReq(url)
	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	if err != nil {
		log.Fatal("Could not parse body", err)
	}
	vv := v.GetArray("matches")
	for i := range vv {
		m = append(m, vv[i].GetInt64("gameId"))
	}
	fmt.Println("Found", len(m), "matches")
	return m
}
