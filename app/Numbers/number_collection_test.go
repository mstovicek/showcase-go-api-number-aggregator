package Numbers

import (
	"reflect"
	"testing"
)

func TestGetUniqueSortedNumbers(t *testing.T) {
	collection := NewNumberCollection()

	collection.AddNumbers([]int{3, 1})
	collection.AddNumbers([]int{4, 2, 1})

	expected := []int{1, 2, 3, 4}
	actual := collection.GetUniqueSortedNumbers()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Numbers are not the same, expected: %v, actual: %v", expected, actual)
	}
}
