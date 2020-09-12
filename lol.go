package main

import (
	"os"
	"strconv"

	"github.com/meximonster/lol-golang/champion"
	"github.com/meximonster/lol-golang/match"
	"github.com/meximonster/lol-golang/player"
)

func main() {
	acc := player.GetaccID(os.Args[1])
	champion := champion.GetChamp(os.Args[2])
	m := player.GetMatches(acc, strconv.Itoa(champion))
	// for i := range m {
	// 	fmt.Println(m[i])
	// }
	match.Info(m[0], champion)
}
