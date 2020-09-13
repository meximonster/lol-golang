package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/meximonster/lol-golang/champion"
	"github.com/meximonster/lol-golang/match"
	"github.com/meximonster/lol-golang/player"
)

func init() {
	if len(os.Args) != 3 {
		fmt.Println("You have to provide exactly a player name and a champion.")
		fmt.Println("Example: go run lol.go xtiltos Thresh")
		os.Exit(1)
	}
}

func main() {
	acc := player.GetaccID(os.Args[1])
	champion := champion.GetChamp(os.Args[2])
	m := player.GetMatches(acc, strconv.Itoa(champion))
	fmt.Println("Found", len(m), "matches")
	statsChan := make(chan match.PlayerStats)
	wg := sync.WaitGroup{}
	go func() {
		for s := range statsChan {
			if s.Timeline.Role != "NONE" {
				fmt.Println("Role:", s.Timeline.Role)
			}
			if s.Timeline.Lane != "NONE" {
				fmt.Println("Lane:", s.Timeline.Lane)
			}
			fmt.Println("K/D/A:", s.Stats.Kills, "/", s.Stats.Deaths, "/", s.Stats.Assists)
			fmt.Println("CS difference per min (0-10):", s.Timeline.CsDiffPerMinDeltas.Zero10, "(10-20):", s.Timeline.CsDiffPerMinDeltas.One020, "(20-30):", s.Timeline.CsDiffPerMinDeltas.Two030)
			fmt.Println("Creeps per min (0-10):", s.Timeline.CreepsPerMinDeltas.Zero10, "(10-20):", s.Timeline.CreepsPerMinDeltas.One020, "(20-30):", s.Timeline.CreepsPerMinDeltas.Two030)
			fmt.Println("Gold per min (0-10):", s.Timeline.GoldPerMinDeltas.Zero10, "(10-20):", s.Timeline.GoldPerMinDeltas.One020, "(10-20):", s.Timeline.GoldPerMinDeltas.Two030)
			fmt.Print("\n\n")
		}
	}()
	for i := range m {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			match.Info(m[index], champion, statsChan)
		}(i)
	}
	wg.Wait()
	close(statsChan)
}
