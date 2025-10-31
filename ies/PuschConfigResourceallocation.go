package ies

import "rrc/utils"

// PUSCH-Config-resourceAllocation ::= ENUMERATED
type PuschConfigResourceallocation struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigResourceallocationEnumeratedNothing = iota
	PuschConfigResourceallocationEnumeratedResourceallocationtype0
	PuschConfigResourceallocationEnumeratedResourceallocationtype1
	PuschConfigResourceallocationEnumeratedDynamicswitch
)
