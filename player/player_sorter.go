package player

import (
	"sort"
)

// By is the type of a "less" function that defines the ordering of its Player arguments.
type By func(p1, p2 Sharer) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(players []Sharer) {
	ps := &playerSorter{
		players: players,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Players to be sorted.
type playerSorter struct {
	players []Sharer
	by      func(p1, p2 Sharer) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *playerSorter) Len() int {
	return len(s.players)
}

// Swap is part of sort.Interface.
func (s *playerSorter) Swap(i, j int) {
	s.players[i], s.players[j] = s.players[j], s.players[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *playerSorter) Less(i, j int) bool {
	return s.by(s.players[i], s.players[j])
}
