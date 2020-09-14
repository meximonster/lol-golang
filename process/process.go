package processdata

import (
	"fmt"
	"sync"

	"github.com/meximonster/lol-golang/match"
)

// Print prints the data
func Print(ch chan match.PlayerStats) {
	for s := range ch {
		if s.Timeline.Role != "NONE" {
			fmt.Println("Role:", s.Timeline.Role)
		}
		if s.Timeline.Lane != "NONE" {
			fmt.Println("Lane:", s.Timeline.Lane)
		}
		fmt.Println("K/D/A:", s.Stats.Kills, "/", s.Stats.Deaths, "/", s.Stats.Assists)
		fmt.Println("CS difference per min (0-10):", s.Timeline.CsDiffPerMinDeltas.Zero10, "(10-20):", s.Timeline.CsDiffPerMinDeltas.One020, "(20-30):", s.Timeline.CsDiffPerMinDeltas.Two030)
		fmt.Println("Creeps per min (0-10):", s.Timeline.CreepsPerMinDeltas.Zero10, "(10-20):", s.Timeline.CreepsPerMinDeltas.One020, "(20-30):", s.Timeline.CreepsPerMinDeltas.Two030)
		fmt.Println("Gold per min (0-10):", s.Timeline.GoldPerMinDeltas.Zero10, "(10-20):", s.Timeline.GoldPerMinDeltas.One020, "(20-30):", s.Timeline.GoldPerMinDeltas.Two030)
		fmt.Print("\n\n")
	}
	close(ch)
}

// Matches calls match.Info which adds stats to channel
func Matches(m []int64, i int, wg *sync.WaitGroup, champion int, ch chan (match.PlayerStats)) {
	defer wg.Done()
	match.Info(m[i], champion, ch)
}
