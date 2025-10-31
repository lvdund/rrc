package ies

import "rrc/utils"

// BandParametersTxSL-r14-v2x-HighPower-r14 ::= ENUMERATED
type BandparameterstxslR14V2xHighpowerR14 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterstxslR14V2xHighpowerR14EnumeratedNothing = iota
	BandparameterstxslR14V2xHighpowerR14EnumeratedSupported
)
