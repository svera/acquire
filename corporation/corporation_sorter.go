package corporation

import (
	"sort"
)

// By is the type of a "less" function that defines the ordering of its Player arguments.
type By func(c1, c2 Interface) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(corporations []Interface) {
	ps := &corporationSorter{
		corporations: corporations,
		by:           by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// corporationSorter joins a By function and a slice of corporations to be sorted.
type corporationSorter struct {
	corporations []Interface
	by           func(c1, c2 Interface) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *corporationSorter) Len() int {
	return len(s.corporations)
}

// Swap is part of sort.Interface.
func (s *corporationSorter) Swap(i, j int) {
	s.corporations[i], s.corporations[j] = s.corporations[j], s.corporations[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *corporationSorter) Less(i, j int) bool {
	return s.by(s.corporations[i], s.corporations[j])
}
