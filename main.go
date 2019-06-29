package main

import (
	"fmt"
	"os"

	"github.com/aarthikrao/matchMaker/teams"
	"github.com/aarthikrao/matchMaker/utils"

	"github.com/aarthikrao/matchMaker/models"
	"github.com/abiosoft/ishell"
)

func main() {
	// A List to store all players
	allPlayers := []models.Player{}

	shell := ishell.New()

	// CLI handling for creating a new user
	shell.AddCmd(&ishell.Cmd{
		Name: "new-player",
		Help: "usage : new-player <name> <rank> | Adds new player to the list",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 2 {
				fmt.Println("Incorrect input, try 'help'")
				return
			}
			val, err := utils.StrToFloat32(c.Args[1])
			if err == nil {
				newPlayer := models.Player{
					Name: c.Args[0],
					Rank: val,
				}
				allPlayers = append(allPlayers, newPlayer)
				fmt.Printf("New Player: %+v \n", newPlayer)
			} else {
				fmt.Println("Please enter a valid number")
			}
		},
	})

	// CLI handling for clearing the player list
	shell.AddCmd(&ishell.Cmd{
		Name: "clear-player-list",
		Help: "usage : clear-player-list | Removes all the players from the list",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 0 {
				fmt.Println("Incorrect input, try 'help'")
				return
			}
			allPlayers = nil
			allPlayers = []models.Player{}
			fmt.Println("Cleared all player list")
		},
	})

	// CLI handling for displaying all the players in the list
	shell.AddCmd(&ishell.Cmd{
		Name: "all-players",
		Help: "usage : all-players | Displays a list of all players",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 0 {
				fmt.Println("Incorrect input, try 'help'")
				return
			}
			for i, player := range allPlayers {
				fmt.Printf("%d %+v \n", i, player)
			}
		},
	})

	// CLI handling for match making by team size
	shell.AddCmd(&ishell.Cmd{
		Name: "make-match",
		Help: "usage : make-match <team-size> | Displays matches that are possible in ascending order of team rank diffrences",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				fmt.Println("Incorrect input, try 'help'")
				return
			}
			teamSize, err := utils.StrToInt(c.Args[0])
			if teamSize > len(allPlayers)/2 {
				fmt.Println("Please enter a value from 0 to", len(allPlayers)/2)
				return
			}
			if err == nil {
				for i, match := range teams.CreateMatches(allPlayers, teamSize) {
					fmt.Printf("%d %v => Diff: %v \n", i, match, utils.Abs(match.TeamA.ARank-match.TeamB.ARank))
				}
			} else {
				fmt.Println("Please enter a valid number")
			}
		},
	})

	// add all program termination related logic here
	shell.Interrupt(func(c *ishell.Context, count int, str string) {
		fmt.Println("Exiting Match maker ...")
		os.Exit(0)
	})

	shell.Run()

}
