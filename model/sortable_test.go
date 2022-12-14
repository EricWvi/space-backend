package model

import (
	"fmt"
	"testing"
)

type Node struct {
	Pos int
	Pre int
}

func (n Node) Prev() int {
	return n.Pre
}

func (n Node) Curr() int {
	return n.Pos
}

func (n Node) IsHead() bool {
	return n.Pre == 0
}

func TestSort(t *testing.T) {
	type args[T comparable] struct {
		elems []ListNode[T]
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []ListNode[T]
	}
	tests := []testCase[int]{
		{
			name: "1",
			args: args[int]{
				elems: []ListNode[int]{
					&Node{
						Pos: 3,
						Pre: 2,
					},
					&Node{
						Pos: 1,
						Pre: 0,
					},
					&Node{
						Pos: 4,
						Pre: 3,
					},
					&Node{
						Pos: 2,
						Pre: 1,
					},
				},
			},
			want: []ListNode[int]{
				&Node{
					Pos: 1,
					Pre: 0,
				},
				&Node{
					Pos: 2,
					Pre: 1,
				},
				&Node{
					Pos: 3,
					Pre: 2,
				},
				&Node{
					Pos: 4,
					Pre: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := Sort(tt.args.elems)
			for _, v := range list {
				node := v.(*Node)
				fmt.Printf("%T: %#v", node, node)
			}
		})
	}
}
