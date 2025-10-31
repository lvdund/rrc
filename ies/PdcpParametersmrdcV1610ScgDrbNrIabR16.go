package ies

import "rrc/utils"

// PDCP-ParametersMRDC-v1610-scg-DRB-NR-IAB-r16 ::= ENUMERATED
type PdcpParametersmrdcV1610ScgDrbNrIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersmrdcV1610ScgDrbNrIabR16EnumeratedNothing = iota
	PdcpParametersmrdcV1610ScgDrbNrIabR16EnumeratedSupported
)
