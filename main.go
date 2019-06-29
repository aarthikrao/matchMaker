package main

import (
	"fmt"

	"github.com/aarthikrao/matchMaker/models"
	"github.com/aarthikrao/matchMaker/teams"
)

func main() {
	ch := make(chan models.EligibleTeams)
	go teams.GetCombinations(ch, 6, 3)
	for i := range ch {
		fmt.Println(i)
	}
}
