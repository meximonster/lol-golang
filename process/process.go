package processdata

import (
	"fmt"
	"sync"

	"github.com/meximonster/lol-golang/match"
)

// Print prints the data
func Print(ch chan match.PlayerStats, wg *sync.WaitGroup) {
	c := 1
	for s := range ch {
		printHelper(s, c)
		c++
		wg.Done()
	}
}

func printHelper(s match.PlayerStats, c int) {
	fmt.Println(c)
	if s.Timeline.Role != "NONE" {
		fmt.Println("Role:", s.Timeline.Role)
	}
	if s.Timeline.Lane != "NONE" {
		fmt.Println("Lane:", s.Timeline.Lane)
	}
	fmt.Println("K/D/A:", s.Stats.Kills, "/", s.Stats.Deaths, "/", s.Stats.Assists)
	//fmt.Println("CS difference per min (0-10):", s.Timeline.CsDiffPerMinDeltas.Zero10, "(10-20):", s.Timeline.CsDiffPerMinDeltas.One020, "(20-30):", s.Timeline.CsDiffPerMinDeltas.Two030)
	//fmt.Println("Creeps per min (0-10):", s.Timeline.CreepsPerMinDeltas.Zero10, "(10-20):", s.Timeline.CreepsPerMinDeltas.One020, "(20-30):", s.Timeline.CreepsPerMinDeltas.Two030)
	//fmt.Println("Gold per min (0-10):", s.Timeline.GoldPerMinDeltas.Zero10, "(10-20):", s.Timeline.GoldPerMinDeltas.One020, "(20-30):", s.Timeline.GoldPerMinDeltas.Two030)
	fmt.Println("")
}

// Matches calls match.Info which adds stats to channel
func Matches(m []int64, i int, champion int, ch chan (match.PlayerStats)) {
	match.Info(m[i], champion, ch)
}
