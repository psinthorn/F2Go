package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 3, 5, 7, 2}

	result := BubbleSort(elements)
	assert.NotNil(t, result)
	assert.EqualValues(t, len(result), 5)
	assert.EqualValues(t, 2, result[0])
	assert.EqualValues(t, 3, result[1])
	assert.EqualValues(t, 5, result[2])
	assert.EqualValues(t, 7, result[3])
	assert.EqualValues(t, 2, result[4])
}

func getElements(n int) []int {
	results := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		results[i] = j
		i++
	}
	return results
}

func benchMarkBubbleSort(b *testing.B) {

}
