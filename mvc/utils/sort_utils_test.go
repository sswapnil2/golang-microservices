package utils

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBubbleSortBestCase(t *testing.T) {
	arr := make([]int, 5)
	for i:= len(arr) - 1; i >= 0; i-- {
		arr[i] = i;
	}
	BubbleSort(arr)
	assert.Less(t, arr[0], arr[1])
}


func getElements(num int) [] int {
	arr := make([]int, num)
	for i:=0; i< num ; i++ {
		arr[i] = rand.Intn(50)
	}
	return arr
}



func BenchmarkBubbleSort(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		BubbleSort(getElements(500))
	}
}
