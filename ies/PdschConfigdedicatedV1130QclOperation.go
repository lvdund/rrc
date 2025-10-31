package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1130-qcl-Operation ::= ENUMERATED
type PdschConfigdedicatedV1130QclOperation struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1130QclOperationEnumeratedNothing = iota
	PdschConfigdedicatedV1130QclOperationEnumeratedTypea
	PdschConfigdedicatedV1130QclOperationEnumeratedTypeb
)
