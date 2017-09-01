package Numbers

import (
	"sort"
)

type NumberCollection interface {
	AddNumbers(numbers []int)
	GetUniqueSortedNumbers() []int
}

type numberCollection struct {
	numbers []int
}

func NewNumberCollection() NumberCollection {
	return newCollectionWithNumbers([]int{})
}

func newCollectionWithNumbers(numbers []int) NumberCollection {
	return &numberCollection{
		numbers: numbers,
	}
}

func (n *numberCollection) AddNumbers(numbers []int) {
	n.numbers = append(n.numbers, numbers...)
}

func (n *numberCollection) GetUniqueSortedNumbers() []int {
	n.sort()
	n.unique()

	return n.numbers
}

func (n *numberCollection) sort() {
	sort.Ints(n.numbers)
}

func (n *numberCollection) unique() {
	seen := make(map[int]struct{}, len(n.numbers))

	i := 0
	for _, num := range n.numbers {
		if _, exists := seen[num]; exists {
			continue
		}

		seen[num] = struct{}{}
		n.numbers[i] = num
		i++
	}

	n.numbers = n.numbers[:i]
}
