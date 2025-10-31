package ies

import "rrc/utils"

// BandParametersRxSL-r14-v2x-HighReception-r14 ::= ENUMERATED
type BandparametersrxslR14V2xHighreceptionR14 struct {
	Value utils.ENUMERATED
}

const (
	BandparametersrxslR14V2xHighreceptionR14EnumeratedNothing = iota
	BandparametersrxslR14V2xHighreceptionR14EnumeratedSupported
)
