package model

type ListNode[T comparable] interface {
	Prev() T
	Curr() T
	IsHead() bool
}

func Sort[T comparable](elems []ListNode[T]) []ListNode[T] {
	links := make(map[T]T)
	values := make(map[T]ListNode[T])
	list := make([]ListNode[T], 0)
	for i := range elems {
		links[elems[i].Prev()] = elems[i].Curr()
		values[elems[i].Curr()] = elems[i]
		if elems[i].IsHead() {
			list = append(list, elems[i])
		}
	}
	if len(list) != 0 {
		curr := list[0]
		for {
			next, ok := links[curr.Curr()]
			if ok {
				v := values[next]
				list = append(list, v)
				curr = v
			} else {
				break
			}
		}
	}
	return list
}
