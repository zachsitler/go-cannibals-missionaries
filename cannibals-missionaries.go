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

// Run BFS on the state space and search for the steps to the solution
func main() {
	Q = lang.NewQueue()
	Q.Push(state{mstart, cstart, start, nil})

	for Q.Len() != 0 {
		v := Q.Poll().(state)

		if goalState(v) {
			fmt.Println("Solution found.")
			fmt.Println("Steps:")
			v.stateHistory()
			break
		}

		v.getChildren()
	}
}

// We only need the parent state for printing out the history of the steps.
// Otherwise, we could do without it.
type state struct {
	m, c, direction int
	parent          *state
}

// Get all possible moves we could make and attempt to add them to the queue.
func (s *state) getChildren() {
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
			if (cstart-s.c) == c && (mstart-s.m) == m {
				continue
			}
			addState(s, c, m)
		}
	}
}

// Ensure that a move doesn't kill any missionaries and doesn't break any rules
func (s *state) validState() bool {
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

// We use -1/1 for the directions so we can easily add and remove
// cannibals/missionaries when moving them.
func (s *state) getDirection() int {
	if s.direction == start {
		return end
	}

	return start
}

func (s *state) stateHistory() {
	path := []string{}
	for s != nil {
		path = append(path, s.printState())
		s = s.parent
	}

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println(path[i])
	}
}

func (s *state) printState() (result string) {
	result = fmt.Sprintf("M: %d, C: %d", s.m, s.c)
	if s.direction == end {
		result += " ----------------<> "
	} else {
		result += " <>---------------- "
	}
	result += fmt.Sprintf("M: %d, C: %d", mstart-s.m, cstart-s.c)
	return
}

// Add a state to the queue if it is valid.
func addState(parent *state, c, m int) {
	direction := parent.getDirection()

	// This is a "move" from one side to the other.
	child := state{parent.m + m*direction, parent.c + c*direction,
		direction, parent}

	if child.validState() {
		Q.Push(child)
	}
}

// This is our stop condition while searching. If we have a state
// that has no cannibals/missionaries and the boat is on the other side
// side then we did it.
func goalState(s state) bool {
	if s.m+s.c != 0 || s.direction != end {
		return false
	}

	return true
}
