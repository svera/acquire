package board

import "github.com/svera/acquire/interfaces"

type sortableCorporations []interfaces.Corporation

func (s sortableCorporations) Len() int           { return len(s) }
func (s sortableCorporations) Less(i, j int) bool { return s[i].Size() < s[j].Size() }
func (s sortableCorporations) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
