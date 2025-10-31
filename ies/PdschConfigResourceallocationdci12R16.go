package ies

import "rrc/utils"

// PDSCH-Config-resourceAllocationDCI-1-2-r16 ::= ENUMERATED
type PdschConfigResourceallocationdci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigResourceallocationdci12R16EnumeratedNothing = iota
	PdschConfigResourceallocationdci12R16EnumeratedResourceallocationtype0
	PdschConfigResourceallocationdci12R16EnumeratedResourceallocationtype1
	PdschConfigResourceallocationdci12R16EnumeratedDynamicswitch
)
