package ies

import "rrc/utils"

// RF-Parameters-v1310-reducedIntNonContComb-r13 ::= ENUMERATED
type RfParametersV1310ReducedintnoncontcombR13 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersV1310ReducedintnoncontcombR13EnumeratedNothing = iota
	RfParametersV1310ReducedintnoncontcombR13EnumeratedSupported
)
