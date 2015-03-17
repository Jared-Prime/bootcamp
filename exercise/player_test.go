package exercise

import (
	"fmt"
	"testing"
)

type player_data struct {
	TestName     string
	UserId       int
	UserName     string
	UserLocation string
	PlayerId     int
}

func TestNewPlayer(t *testing.T) {
	tests := []player_data{{"Basic Player test", 42, "Jared", "Chicago", 1234}}

	for _, test := range tests {
		NewPlayer(test.UserId, test.UserName, test.UserLocation, test.PlayerId)
	}
}

func TestGreetings(t *testing.T) {
	var expected string
	tests := []player_data{{"Basic Player test", 42, "Jared", "Chicago", 1234}}

	for _, test := range tests {
		player := NewPlayer(test.UserId, test.UserName, test.UserLocation, test.PlayerId)
		expected = fmt.Sprintf("Hi %s from %s", test.UserName, test.UserLocation)
		if result := player.Greetings(); result != expected {
			t.Errorf("%s: expected %s, got %s", test.TestName, expected, player.Greetings())
		}
	}
}
