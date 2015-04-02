package exercise

import (
	"testing"
)

var (
	users = []string{
		"Matthew", "Peter", "Giana", "Adriano", "Elizabeth", "Sarah", "Augustus",
		"Heidi", "Emilie", "Aaron"}
	coinCount    = 50
	expectedDist = map[string]int{
		"Matthew": 2, "Peter": 2, "Giana": 4, "Adriano": 7,
		"Elizabeth": 5, "Sarah": 2, "Augustus": 10, "Heidi": 5,
		"Emilie": 6, "Aaron": 5}
	expectedLeft = 2
)

func TestDistributeBitcoins(t *testing.T) {
	result, actualLeft := DistributeBitcoins(users, coinCount)
	if actualLeft != expectedLeft {
		t.Errorf("expected %d coins left, got %d", expectedLeft, actualLeft)
	}

	for user, coins := range expectedDist {
		if actual := result[user]; actual != coins {
			t.Errorf("Expected %s to have %d coins, got: %d", user, coins, actual)
		}
	}
}
