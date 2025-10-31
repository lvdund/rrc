package ies

import "rrc/utils"

// RF-Parameters-extendedBand-n77-r16 ::= ENUMERATED
type RfParametersExtendedbandN77R16 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersExtendedbandN77R16EnumeratedNothing = iota
	RfParametersExtendedbandN77R16EnumeratedSupported
)
