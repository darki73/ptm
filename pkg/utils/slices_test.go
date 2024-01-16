package utils

import (
	"reflect"
	"testing"
)

// TestRemoveEmptyStringsFromSlice tests the RemoveEmptyStringsFromSlice function.
func TestRemoveEmptyStringsFromSlice(t *testing.T) {
	input1 := []string{"1", "2", "", "3", ""}
	expected1 := []string{"1", "2", "3"}
	output1 := RemoveEmptyStringsFromSlice(input1)

	if !reflect.DeepEqual(output1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, output1)
	}

	input2 := []string{"", "", ""}
	var expected2 []string
	output2 := RemoveEmptyStringsFromSlice(input2)

	if !reflect.DeepEqual(output2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, output2)
	}

	input3 := []string{"1", "2", "3"}
	expected3 := []string{"1", "2", "3"}
	output3 := RemoveEmptyStringsFromSlice(input3)

	if !reflect.DeepEqual(output3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, output3)
	}
}
