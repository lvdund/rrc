package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1530-ce-PDSCH-FlexibleStartPRB-AllocConfig-r15 ::= ENUMERATED
type PdschConfigdedicatedV1530CePdschFlexiblestartprbAllocconfigR15 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1530CePdschFlexiblestartprbAllocconfigR15EnumeratedNothing = iota
	PdschConfigdedicatedV1530CePdschFlexiblestartprbAllocconfigR15EnumeratedOn
)
