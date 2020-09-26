package main

import (
	"fmt"
	"lol-golang/champion"
	"lol-golang/match"
	"lol-golang/player"
	processdata "lol-golang/process"
	"os"
	"strconv"
	"sync"
	"time"
)

func init() {
	if len(os.Args) != 3 {
		fmt.Println("You have to provide exactly a player name and a champion.")
		fmt.Println("Example: go run lol.go xtiltos Thresh")
		os.Exit(1)
	}
}

func main() {
	start := time.Now()
	numWorkers := 20

	acc := player.GetaccID(os.Args[1])
	champion := champion.GetChamp(os.Args[2])
	m := player.GetMatches(acc, strconv.Itoa(champion))
	matchIds := make(chan int64)
	go func() {
		for _, id := range m {
			matchIds <- id
		}
		close(matchIds)
	}()

	results := make(chan match.PlayerStats)
	wg := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		wg2.Add(1)
		go func() {
			defer wg.Done()
			defer wg2.Done()
			match.Info(champion, matchIds, results)
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		processdata.Print(results, len(m))
	}()
	wg2.Wait()
	close(results)
	wg.Wait()
	fmt.Println("----------------")
	fmt.Println("Finished in", time.Since(start))
}
