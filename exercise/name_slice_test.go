package exercise

import (
	"testing"
)

var (
	inputSlice = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
		"Emma", "Isabella", "Emily", "Madison", "Ava", "Olivia", "Sophia", "Abigail",
		"Elizabeth", "Chloe", "Samantha", "Addison", "Natalie", "Mia", "Alexis"}
	expectedOutputSlice = [][]string{
		[]string{},
		[]string{},
		[]string{"Ava", "Mia"},
		[]string{"Evan", "Neil", "Adam", "Matt", "Emma"},
		[]string{"Emily", "Chloe"},
		[]string{"Martin", "Olivia", "Sophia", "Alexis"},
		[]string{"Katrina", "Madison", "Abigail", "Addison", "Natalie"},
		[]string{"Isabella", "Samantha"},
		[]string{"Elizabeth"},
	}
	expectedSize = len(expectedOutputSlice)
)

func TestOrganizeNames(t *testing.T) {
	resultSlice := OrganizeNames(inputSlice)
	if size := len(resultSlice); size < expectedSize {
		t.Errorf("got: %d slices\nexpected: %d", size, expectedSize)
	}

	// @TODO: improve the set comparison, as this is order dependent
	for i, result := range resultSlice {
		for j, resultString := range result {
			if expectedString := expectedOutputSlice[i][j]; expectedString != resultString {
				t.Errorf("%s does not match %s", resultString, expectedString)
			}
		}
	}
}
