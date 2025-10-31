package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1530-qcl-Operation-v1530 ::= ENUMERATED
type PdschConfigdedicatedV1530QclOperationV1530 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1530QclOperationV1530EnumeratedNothing = iota
	PdschConfigdedicatedV1530QclOperationV1530EnumeratedTypec
)
