package ies

import "rrc/utils"

// MRDC-Parameters-tdm-Pattern ::= ENUMERATED
type MrdcParametersTdmPattern struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersTdmPatternEnumeratedNothing = iota
	MrdcParametersTdmPatternEnumeratedSupported
)
