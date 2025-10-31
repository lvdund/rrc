package ies

import "rrc/utils"

// SRS-Config-tpc-Accumulation ::= ENUMERATED
type SrsConfigTpcAccumulation struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigTpcAccumulationEnumeratedNothing = iota
	SrsConfigTpcAccumulationEnumeratedDisabled
)
