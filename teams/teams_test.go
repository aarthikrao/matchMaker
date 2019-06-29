package teams

import (
	"testing"

	"github.com/aarthikrao/matchMaker/models"
)

func TestGetCombinations(t *testing.T) {
	type args struct {
		n     int
		k     int
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Positive", args{5, 2, 15}},
		{"Positive", args{6, 3, 10}},
	}
	for _, tt := range tests {
		ch := make(chan models.EligibleTeams)
		t.Run(tt.name, func(t *testing.T) {
			go GetCombinations(ch, tt.args.n, tt.args.k)
			count := 0
			for _ = range ch {
				// We are only checking the count of matches.
				count++
			}
			if count != tt.args.count {
				t.Errorf("Abs() = %v, want %v", count, tt.args.count)
			}
		})
	}
}
