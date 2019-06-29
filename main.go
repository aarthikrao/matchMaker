package main

import (
	"fmt"
	"sort"

	"github.com/aarthikrao/matchMaker/models"
	"github.com/aarthikrao/matchMaker/teams"
)

func main() {
	// Create new members
	listOfMembers := []models.Member{}
	for i := 0; i < 6; i++ {
		newMember := models.Member{
			Name: fmt.Sprintf("player%d", i),
			Rank: i,
		}
		listOfMembers = append(listOfMembers, newMember)
	}
	fmt.Println(listOfMembers)

	allMatches := []models.Match{}
	ch := make(chan models.EligibleTeams)
	go teams.GetCombinations(ch, len(listOfMembers), 2)
	for eT := range ch {
		newMatch := models.Match{}
		teamAAvg := 0
		for _, memberInTeamA := range eT.TeamA {
			newMatch.TeamA.Members = append(newMatch.TeamA.Members, listOfMembers[memberInTeamA])
			teamAAvg += listOfMembers[memberInTeamA].Rank
		}
		teamAAvg = teamAAvg / len(newMatch.TeamA.Members)
		newMatch.TeamA.ARank = teamAAvg
		teamBAvg := 0
		for _, memberInTeamB := range eT.TeamB {
			newMatch.TeamB.Members = append(newMatch.TeamB.Members, listOfMembers[memberInTeamB])
			teamBAvg += listOfMembers[memberInTeamB].Rank
		}
		teamBAvg = teamBAvg / len(newMatch.TeamB.Members)
		newMatch.TeamB.ARank = teamBAvg
		allMatches = append(allMatches, newMatch)
		fmt.Println(newMatch)
	}
	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i].TeamA.ARank-allMatches[i].TeamB.ARank > allMatches[j].TeamA.ARank-allMatches[j].TeamB.ARank
	})
	fmt.Println("######################################")
	for _, match := range allMatches {
		x := match.TeamA.ARank - match.TeamB.ARank
		fmt.Println(x)
	}
}
