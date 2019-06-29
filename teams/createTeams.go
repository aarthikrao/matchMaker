package teams

import (
	"fmt"

	"github.com/aarthikrao/matchMaker/models"

	"github.com/gonum/stat/combin"
)

// GetCombinations generates a combination of 2 k teams from n members.
// Channel ch is used to push the teams
func GetCombinations(ch chan models.EligibleTeams, n int, k int) {
	teamCombinations := combin.Combinations(n, k)
	fmt.Println(teamCombinations)
	for i := 0; i < len(teamCombinations)-1; i++ {
		for j := i + 1; j < len(teamCombinations); j++ {
			if duplicateMemberExsists(teamCombinations[j], teamCombinations[i]) {
				eT := models.EligibleTeams{
					TeamA: teamCombinations[i],
					TeamB: teamCombinations[j],
				}
				ch <- eT
			}
		}
	}
	close(ch)
}

func duplicateMemberExsists(teamA []int, teamB []int) bool {
	for _, memberTeamA := range teamA {
		for _, memberTeamB := range teamB {
			if memberTeamA == memberTeamB {
				return false
			}
		}
	}
	return true
}
