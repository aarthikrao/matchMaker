package models

type EligibleTeams struct {
	TeamA []int
	TeamB []int
}

type Member struct {
	Name string
	Rank int
}

type Team struct {
	Members []Member
	ARank   int // Average rank of the team
}

type Match struct {
	TeamA Team
	TeamB Team
}
