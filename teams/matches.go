package teams

import (
	"sort"

	"github.com/aarthikrao/matchMaker/utils"

	"github.com/aarthikrao/matchMaker/models"
)

// CreateMatches is a util function to create matches from a list of players and team size
func CreateMatches(allPlayers []models.Player, teamSize int) []models.Match {
	allMatches := []models.Match{}

	// Create a channel for fetching the eligible teams
	matchedPlayersChannel := make(chan models.EligibleTeams)

	// GetCombinations will run on a seperate go-routine passing the generated matches by matchedPlayersChannel
	go GetCombinations(matchedPlayersChannel, len(allPlayers), teamSize)

	for eT := range matchedPlayersChannel {
		newMatch := models.Match{}

		// Map all the players of team A
		teamAAvg := float32(0)
		for _, memberInTeamA := range eT.TeamA {
			newMatch.TeamA.Players = append(newMatch.TeamA.Players, allPlayers[memberInTeamA])
			teamAAvg += allPlayers[memberInTeamA].Rank
		}
		// Calculate the average rank of team and assign to team avg
		teamAAvg = teamAAvg / float32(len(newMatch.TeamA.Players))
		newMatch.TeamA.ARank = teamAAvg

		// Map all the players of team B
		teamBAvg := float32(0)
		for _, memberInTeamB := range eT.TeamB {
			newMatch.TeamB.Players = append(newMatch.TeamB.Players, allPlayers[memberInTeamB])
			teamBAvg += float32(allPlayers[memberInTeamB].Rank)
		}
		// Calculate the average rank of team and assign to team avg
		teamBAvg = teamBAvg / float32(len(newMatch.TeamB.Players))
		newMatch.TeamB.ARank = teamBAvg

		// Add match to allMatches list
		allMatches = append(allMatches, newMatch)
	}
	sortMatchesByScoreDiffrence(allMatches)
	return allMatches
}

// sortMatchesByScoreDiffrence sorts all the matches by their average score diffrence
// write all the sorting logic here
func sortMatchesByScoreDiffrence(allMatches []models.Match) {
	sort.Slice(allMatches, func(i, j int) bool {
		return utils.Abs(allMatches[i].TeamA.ARank-allMatches[i].TeamB.ARank) < utils.Abs(allMatches[j].TeamA.ARank-allMatches[j].TeamB.ARank)
	})
}
