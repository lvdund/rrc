package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-p-a ::= ENUMERATED
type PdschConfigdedicatedPA struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedPAEnumeratedNothing = iota
	PdschConfigdedicatedPAEnumeratedDb_6
	PdschConfigdedicatedPAEnumeratedDb_4dot77
	PdschConfigdedicatedPAEnumeratedDb_3
	PdschConfigdedicatedPAEnumeratedDb_1dot77
	PdschConfigdedicatedPAEnumeratedDb0
	PdschConfigdedicatedPAEnumeratedDb1
	PdschConfigdedicatedPAEnumeratedDb2
	PdschConfigdedicatedPAEnumeratedDb3
)
