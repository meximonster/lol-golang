package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/meximonster/lol-golang/champion"
	"github.com/meximonster/lol-golang/match"
	"github.com/meximonster/lol-golang/player"
	processdata "github.com/meximonster/lol-golang/process"
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
	statsChan := make(chan match.PlayerStats)
	wg := &sync.WaitGroup{}
	go processdata.Print(statsChan, wg)
	for i := range m {
		if i%19 == 0 {
			time.Sleep(1 * time.Second)
		}
		wg.Add(1)
		go processdata.Matches(m, i, champion, statsChan)

	}
	wg.Wait()
	close(statsChan)
}
