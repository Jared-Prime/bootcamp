package exercise

func DistributeBitcoins(users []string, coins int) (map[string]int, int) {
	userCount := make(map[string]int)

	for _, user := range users {
		for _, char := range user {
			for award := awardCoins(char); award > 0; award-- {
				if coins <= 0 || userCount[user] >= 10 {
					break
				}

				userCount[user]++
				coins--
			}
		}
	}

	return userCount, coins
}

func awardCoins(char rune) int {
	switch char {
	case 'A', 'E', 'a', 'e':
		return 1
	case 'I', 'i':
		return 2
	case 'O', 'o':
		return 3
	case 'U', 'u':
		return 4
	default:
		return 0
	}
}
