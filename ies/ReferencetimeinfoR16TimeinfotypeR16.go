package ies

import "rrc/utils"

// ReferenceTimeInfo-r16-timeInfoType-r16 ::= ENUMERATED
type ReferencetimeinfoR16TimeinfotypeR16 struct {
	Value utils.ENUMERATED
}

const (
	ReferencetimeinfoR16TimeinfotypeR16EnumeratedNothing = iota
	ReferencetimeinfoR16TimeinfotypeR16EnumeratedLocalclock
)
