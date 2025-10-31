package ies

import "rrc/utils"

// RF-Parameters-extendedBand-n77-2-r17 ::= ENUMERATED
type RfParametersExtendedbandN772R17 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersExtendedbandN772R17EnumeratedNothing = iota
	RfParametersExtendedbandN772R17EnumeratedSupported
)
