package ies

import "rrc/utils"

// PUSCH-Config-transformPrecoder ::= ENUMERATED
type PuschConfigTransformprecoder struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigTransformprecoderEnumeratedNothing = iota
	PuschConfigTransformprecoderEnumeratedEnabled
	PuschConfigTransformprecoderEnumeratedDisabled
)
