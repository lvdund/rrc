package ies

import "rrc/utils"

// RLC-Parameters-extendedT-StatusProhibit-r16 ::= ENUMERATED
type RlcParametersExtendedtStatusprohibitR16 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersExtendedtStatusprohibitR16EnumeratedNothing = iota
	RlcParametersExtendedtStatusprohibitR16EnumeratedSupported
)
