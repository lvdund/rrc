package ies

import "rrc/utils"

// TimeReferenceInfo-r15-timeInfoType-r15 ::= ENUMERATED
type TimereferenceinfoR15TimeinfotypeR15 struct {
	Value utils.ENUMERATED
}

const (
	TimereferenceinfoR15TimeinfotypeR15EnumeratedNothing = iota
	TimereferenceinfoR15TimeinfotypeR15EnumeratedLocalclock
)
