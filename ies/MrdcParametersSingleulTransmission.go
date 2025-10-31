package ies

import "rrc/utils"

// MRDC-Parameters-singleUL-Transmission ::= ENUMERATED
type MrdcParametersSingleulTransmission struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersSingleulTransmissionEnumeratedNothing = iota
	MrdcParametersSingleulTransmissionEnumeratedSupported
)
