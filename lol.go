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
	resultsChan := make(chan match.Stats)
	wg := sync.WaitGroup{}
	go func() {
		for res := range resultsChan {
			fmt.Println(res.Kills, res.Deaths, res.Assists)
		}
	}()
	for i := range m {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			match.Info(m[index], champion, resultsChan)
		}(i)
	}
	wg.Wait()
	close(resultsChan)

}
