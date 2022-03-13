package dp

import (
	"fmt"
	"testing"
)

func TestHowSum(t *testing.T) {
	cache := make(map[int][]int)
	result := HowSum(38, []int{3, 4, 8}, cache)
	fmt.Println(result)
}
