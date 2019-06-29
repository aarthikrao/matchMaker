package teams

import (
	"reflect"
	"testing"

	"github.com/aarthikrao/matchMaker/models"
)

func TestCreateMatches(t *testing.T) {

	// Create all the mock test data
	allPlayer := []models.Player{}
	player1 := models.Player{Name: "player1", Rank: 10}
	player2 := models.Player{Name: "player2", Rank: 20}
	player3 := models.Player{Name: "player3", Rank: 30}
	player4 := models.Player{Name: "player4", Rank: 40}

	allPlayer = append(allPlayer, player1)
	allPlayer = append(allPlayer, player2)
	allPlayer = append(allPlayer, player3)
	allPlayer = append(allPlayer, player4)

	team14 := models.Team{Players: []models.Player{player1, player4}, ARank: 25}
	team23 := models.Team{Players: []models.Player{player2, player3}, ARank: 25}
	team13 := models.Team{Players: []models.Player{player1, player3}, ARank: 20}
	team24 := models.Team{Players: []models.Player{player2, player4}, ARank: 30}
	team12 := models.Team{Players: []models.Player{player1, player2}, ARank: 15}
	team34 := models.Team{Players: []models.Player{player3, player4}, ARank: 35}

	allMatches := []models.Match{}
	allMatches = append(allMatches, models.Match{TeamA: team14, TeamB: team23})
	allMatches = append(allMatches, models.Match{TeamA: team13, TeamB: team24})
	allMatches = append(allMatches, models.Match{TeamA: team12, TeamB: team34})

	type args struct {
		allPlayers []models.Player
		teamSize   int
	}
	tests := []struct {
		name string
		args args
		want []models.Match
	}{
		{"Positive case", args{allPlayers: allPlayer, teamSize: 2}, allMatches},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMatches(tt.args.allPlayers, tt.args.teamSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
