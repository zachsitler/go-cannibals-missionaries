package main

import (
	"fmt"

	"github.com/hishboy/gocommons/lang"
)

const (
	size   = 2
	cstart = 3
	mstart = 3
	left   = 1
	right  = -1
)

var Q *lang.Queue

type state struct {
	m, c, direction int
}

func main() {
	Q = lang.NewQueue()
	Q.Push(state{3, 3, left})
	// iter := 0
	for Q.Len() != 0 {

		// if iter == 5 {
		// 	break
		// }
		// Dequeue the first element
		v := Q.Poll().(state)
		fmt.Println(v)
		if goalState(v) {
			fmt.Println("done!")
			break
		}

		// Generate the adjacent paths
		getSuccessors(v)
		// iter++
	}
}

func getSuccessors(parent state) {
	for c := 0; c <= size; c++ {
		for m := 0; m <= size; m++ {
			// skip the useless case
			if c == 0 && m == 0 {
				continue
			}
			if c+m > size {
				break
			}

			// We don't just want to keep undoing what we just did.
			//  For example, 1 0 and then 0 1 doesn't accomplish
			// anything
			if (cstart-parent.c) == c && (mstart-parent.m) == m {
				continue
			}
			addState(parent, c, m)
		}
	}
}

func addState(parent state, c, m int) {
	direction := getDirection(parent.direction)
	child := state{parent.m + m*direction, parent.c + c*direction, direction}

	if validState(child) {
		Q.Push(child)
	}
}

func getDirection(direction int) int {
	if direction == left {
		return right
	}

	return left
}

func validState(s state) bool {
	if s.m > mstart || s.c > cstart || s.m < 0 || s.c < 0 {
		return false
	}
	if s.m < s.c && s.m > 0 {
		return false
	}
	if (mstart-s.m) < (cstart-s.c) && (mstart-s.m) > 0 {
		return false
	}

	return true
}

func goalState(s state) bool {
	if s.m+s.c != 0 || s.direction == left {
		return false
	}

	return true
}
