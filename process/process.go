package processdata

import (
	"fmt"
	"lol-golang/match"
	"math"
)

// GatheredStats are the unified stats
type GatheredStats struct {
	kills             int
	deaths            int
	assists           int
	damageToChampions int
	goldEarned        int
	visionScore       int
	minionsKilled     int
	fbCounter         int
	fbAssistCounter   int
}

// Print prints the data
func Print(ch chan match.PlayerStats, n int) {
	var kills, deaths, assists, damageToChampions, goldEarned, visionScore, minionsKilled, fbCounter, fbAssistCounter int
	for s := range ch {
		//fmt.Println("reading from channel")
		kills += s.Stats.Kills
		deaths += s.Stats.Deaths
		assists += s.Stats.Assists
		damageToChampions += s.Stats.TotalDamageDealtToChampions
		goldEarned += s.Stats.GoldEarned
		visionScore += s.Stats.VisionScore
		minionsKilled += s.Stats.TotalMinionsKilled
		if s.Stats.FirstBloodKill {
			fbCounter++
		}
	}
	finalStats := GatheredStats{
		kills, deaths, assists, damageToChampions, goldEarned, visionScore, minionsKilled, fbCounter, fbAssistCounter,
	}
	printHelper(finalStats, n)
}

// PrintHelper helps
func printHelper(s GatheredStats, n int) {
	fmt.Println("Average KDA:", s.kills/n, "/", s.deaths/n, "/", s.assists/n)
	fmt.Println("KILLS")
	fmt.Println("Total:", s.kills)
	fmt.Println("DEATHS")
	fmt.Println("Total:", s.deaths)
	fmt.Println("ASSISTS")
	fmt.Println("Total:", s.assists)
	fmt.Println("----------------")
	fmt.Println("Total damage to champions:", s.damageToChampions, "Average:", s.damageToChampions/n)
	fmt.Println("Total minions killed:", s.minionsKilled, "Average:", s.minionsKilled/n)
	fmt.Println("Total gold earned:", s.goldEarned, "Average:", s.goldEarned/n)
	fmt.Println("Total vision score:", s.visionScore, "Average:", s.visionScore/n)
	fmt.Println("Total first bloods:", s.fbCounter, "First blood ratio:", math.Round((float64(s.fbCounter)/float64(n))*100), "%")
}

//// Matches calls match.Info which adds stats to channel
//func Matches(m []int64, i int, champion int, ch chan match.PlayerStats, wg *sync.WaitGroup) {
//	defer wg.Done()
//	match.Info(m[i], champion, ch)
//}
//
//// Wait blocks and will eventually close the channel
//func Wait(ch chan match.PlayerStats, wg *sync.WaitGroup) {
//	wg.Wait()
//	//close(ch)
//}
