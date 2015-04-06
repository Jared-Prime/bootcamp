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
	expectedSize        = len(expectedOutputSlice)
	expectedSubsetSizes = []int{0, 0, 2, 5, 2, 4, 5, 2, 1}
)

func TestOrganizeNames(t *testing.T) {
	resultSlice := OrganizeNames(inputSlice)

	if actualSize := len(resultSlice); actualSize != expectedSize {
		t.Errorf("expected result set of size %d, got: %d", expectedSize, actualSize)
	}

	for i, subset := range expectedOutputSlice {
		if actualSize := len(subset); actualSize != expectedSubsetSizes[i] {
			t.Errorf("expected result subset of size %d, got: %d", expectedSubsetSizes[i], actualSize)
		}

		for j, name := range resultSlice[i] {
			if expected := subset[j]; expected != name {
				t.Errorf("expected: %s, got %s", expected, name)
			}
		}
	}
}
