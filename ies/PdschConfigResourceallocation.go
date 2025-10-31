package ies

import "rrc/utils"

// PDSCH-Config-resourceAllocation ::= ENUMERATED
type PdschConfigResourceallocation struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigResourceallocationEnumeratedNothing = iota
	PdschConfigResourceallocationEnumeratedResourceallocationtype0
	PdschConfigResourceallocationEnumeratedResourceallocationtype1
	PdschConfigResourceallocationEnumeratedDynamicswitch
)
