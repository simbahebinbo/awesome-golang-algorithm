package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *Node
		expect [][]int
	}{
		{"TestCase1", &Node{Val: 1}, [][]int{{1}}},
		{"TestCase2", &Node{
			Val: 1,
			Children: []*Node{
				{
					Val: 3,
					Children: []*Node{
						{Val: 5},
						{Val: 6},
					},
				},
				{Val: 2},
				{Val: 4},
			},
		}, [][]int{{1}, {3, 2, 4}, {5, 6}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
