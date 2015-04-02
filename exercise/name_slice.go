package exercise

func OrganizeNames(names []string) [][]string {
	maxLength := getMaxLength(names)
	resultSet := make([][]string, maxLength)

	for _, name := range names {
		curPosition := resultSet[len(name)-1]
		curPosition = append(curPosition, name)
	}
	return resultSet
}

func getMaxLength(names []string) int {
	var maxLength int
	for _, name := range names {
		if length := len(name); length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
