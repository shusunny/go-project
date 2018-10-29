package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{1, 9, 10, 11, 12, 8, 500}, 10},
		test{[]int{1, 19, 20, 21, 500}, 20},
		test{[]int{10, 20, 30, 40, 50, 80, 500}, 44},
		test{[]int{1, 5, 6, 7, 9, 9, 500}, 7.2},
	}

	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Error("got", f, "expect", v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 9, 10, 11, 12, 8, 500}))
	// output:
	// 10
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 9, 10, 11, 12, 13, 8, 7, 200})
	}
}
