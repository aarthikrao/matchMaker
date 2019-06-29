package models

// EligibleTeams defines the combinations of matched teams from the matching algorithm
type EligibleTeams struct {
	TeamA []int
	TeamB []int
}

// Player defines every player in the system
type Player struct {
	Name string
	Rank float32
}

// Team defines a group of players and their average score
type Team struct {
	Players []Player
	ARank   float32 // Average rank of the team
}

// Match defines the two teams that will participate in a match
type Match struct {
	TeamA Team
	TeamB Team
}
