package main

import (
	"fmt"
	"os"

	"github.com/meximonster/lol-golang/champion"
	"github.com/meximonster/lol-golang/player"
)

func main() {
	acc := player.GetaccID(os.Args[1])
	champion := champion.GetChamp(os.Args[2])
	fmt.Println(acc, champion)
}
