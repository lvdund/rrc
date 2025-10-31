package ies

import "rrc/utils"

// PUSCH-Config-resourceAllocationDCI-0-2-r16 ::= ENUMERATED
type PuschConfigResourceallocationdci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigResourceallocationdci02R16EnumeratedNothing = iota
	PuschConfigResourceallocationdci02R16EnumeratedResourceallocationtype0
	PuschConfigResourceallocationdci02R16EnumeratedResourceallocationtype1
	PuschConfigResourceallocationdci02R16EnumeratedDynamicswitch
)
