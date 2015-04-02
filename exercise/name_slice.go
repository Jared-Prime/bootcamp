package exercise

func OrganizeNames(names []string) [][]string {
	var result [][]string
	var pos, curSize int

	for _, name := range names {
		pos = len(name)
		curSize = len(result)
		if curSize < pos {
			newResult := make([][]string, pos)
			copy(newResult, result)
			result = newResult
		}
		result[pos-1] = append(result[pos-1], name)
	}
	return result
}
