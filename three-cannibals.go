package main

import (
	"fmt"

	"github.com/hishboy/gocommons/lang"
)

const (
	size   = 2
	cstart = 3
	mstart = 3
	start  = 1
	end    = -1
)

var Q *lang.Queue

type state struct {
	m, c, direction int
	parent          *state
}

func main() {
	Q = lang.NewQueue()
	Q.Push(state{cstart, mstart, start, nil})
	for Q.Len() != 0 {
		v := Q.Poll().(state)

		// Uncomment to see nodes that were parsed
		// fmt.Println(v)

		if goalState(v) {
			fmt.Println("Solution found.")
			fmt.Println("Steps:")
			stateHistory(v)
			break
		}

		getChildren(v)
	}
}

// Get all possible moves we could make and attempt to add them to the queue.
func getChildren(parent state) {
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
			// For example, 1 0 and then 0 1 doesn't accomplish
			// anything and we will be stuck in an infinite loop.
			if (cstart-parent.c) == c && (mstart-parent.m) == m {
				continue
			}
			addState(parent, c, m)
		}
	}
}

// Add a state to the queue if it is valid.
func addState(parent state, c, m int) {
	direction := getDirection(parent.direction)

	// This is a "move" from one side to the other
	child := state{parent.m + m*direction, parent.c + c*direction,
		direction, &parent}

	if validState(child) {
		Q.Push(child)
	}
}

// Ensure that a move doesn't kill any missionaries and doesn't break any rules
func validState(s state) bool {
	// Make sure we aren't breaking any obvious rules
	if s.m > mstart || s.c > cstart || s.m < 0 || s.c < 0 {
		return false
	}
	// We don't want the cannibals to eat the missionaries
	if s.m < s.c && s.m > 0 || mstart-s.m < cstart-s.c && mstart-s.m > 0 {
		return false
	}

	return true
}

// This is our stop condition while searching. If we have a state
// that has no cannibals/missionaries and the boat is on the other side
// side then we did it.
func goalState(s state) bool {
	if s.m+s.c != 0 || s.direction == start {
		return false
	}

	return true
}

func getDirection(direction int) int {
	if direction == start {
		return end
	}

	return start
}

// Horrible way of printing the history of the final state in reverse
func stateHistory(s state) {
	current := s
	path := []string{}
	for current.parent != nil {
		curPath := fmt.Sprintf("M: %d, C: %d, direction: %d", current.m, current.c, current.direction)
		path = append(path, curPath)
		current = *current.parent
	}

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println(path[i])
	}
}
