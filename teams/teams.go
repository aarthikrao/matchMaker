package teams

import (
	"github.com/aarthikrao/matchMaker/models"

	"github.com/gonum/stat/combin"
)

// GetCombinations generates a combination of 2 k teams from n members.
// Channel ch is used to push the teams
func GetCombinations(ch chan models.EligibleTeams, n int, k int) {
	// Get all possible team combinations of team size k from n players
	teamCombinations := combin.Combinations(n, k)

	// Pass only those teams to the channel where the player doesnt play in both the teams
	for i := 0; i < len(teamCombinations)-1; i++ {
		for j := i + 1; j < len(teamCombinations); j++ {
			if noDuplicateMemberExsists(teamCombinations[j], teamCombinations[i]) {
				eT := models.EligibleTeams{
					TeamA: teamCombinations[i],
					TeamB: teamCombinations[j],
				}

				// pass the match through the channel to the main go-routine
				ch <- eT
			}
		}
	}
	// Close the channel here to indicate to the main go-routine that all the match combinations are done
	close(ch)
}

// returns true only if a player from teamA is not present in teamB
func noDuplicateMemberExsists(teamA []int, teamB []int) bool {
	for _, memberTeamA := range teamA {
		for _, memberTeamB := range teamB {
			if memberTeamA == memberTeamB {
				return false
			}
		}
	}
	return true
}
