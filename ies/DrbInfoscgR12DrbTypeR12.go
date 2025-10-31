package ies

import "rrc/utils"

// DRB-InfoSCG-r12-drb-Type-r12 ::= ENUMERATED
type DrbInfoscgR12DrbTypeR12 struct {
	Value utils.ENUMERATED
}

const (
	DrbInfoscgR12DrbTypeR12EnumeratedNothing = iota
	DrbInfoscgR12DrbTypeR12EnumeratedSplit
	DrbInfoscgR12DrbTypeR12EnumeratedScg
)
