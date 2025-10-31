package ies

import "rrc/utils"

// PagingWeight-NB-r14 ::= ENUMERATED
type PagingweightNbR14 struct {
	Value utils.ENUMERATED
}

const (
	PagingweightNbR14EnumeratedNothing = iota
	PagingweightNbR14EnumeratedW1
	PagingweightNbR14EnumeratedW2
	PagingweightNbR14EnumeratedW3
	PagingweightNbR14EnumeratedW4
	PagingweightNbR14EnumeratedW5
	PagingweightNbR14EnumeratedW6
	PagingweightNbR14EnumeratedW7
	PagingweightNbR14EnumeratedW8
	PagingweightNbR14EnumeratedW9
	PagingweightNbR14EnumeratedW10
	PagingweightNbR14EnumeratedW11
	PagingweightNbR14EnumeratedW12
	PagingweightNbR14EnumeratedW13
	PagingweightNbR14EnumeratedW14
	PagingweightNbR14EnumeratedW15
	PagingweightNbR14EnumeratedW16
)
