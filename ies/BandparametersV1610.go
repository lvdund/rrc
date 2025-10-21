package ies

import "rrc/utils"

// BandParameters-v1610 ::= SEQUENCE
type BandparametersV1610 struct {
	IntrafreqdapsR16                  *interface{}
	AddsrsFrequencyhoppingR16         *utils.ENUMERATED
	AddsrsAntennaswitchingR16         *interface{}
	SrsCapabilityperbandpairlistV1610 *interface{}
}
