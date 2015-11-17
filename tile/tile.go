package tile

import (
	"github.com/svera/acquire/corporation"
)

type Empty struct {
	number int
	letter string
}

func NewEmpty(number int, letter string) *Empty {
	return &Empty{number, letter}
}

func (e *Empty) Number() int {
	return e.number
}

func (e *Empty) Letter() string {
	return e.letter
}

func (e *Empty) ContentType() string {
	return "empty"
}

type Orphan struct {
	number int
	letter string
}

func NewOrphan(number int, letter string) *Orphan {
	return &Orphan{number, letter}
}

func (t *Orphan) Number() int {
	return t.number
}

func (t *Orphan) Letter() string {
	return t.letter
}

func (t *Orphan) ContentType() string {
	return "orphan"
}

type Corporation struct {
	number      int
	letter      string
	corporation corporation.Interface
}

func NewCorporation(number int, letter string, corporation corporation.Interface) *Corporation {
	return &Corporation{number, letter, corporation}
}

func (t *Corporation) Number() int {
	return t.number
}

func (t *Corporation) Letter() string {
	return t.letter
}

func (t *Corporation) ContentType() string {
	return "corporation"
}

func (t *Corporation) Corporation() corporation.Interface {
	return t.corporation
}
