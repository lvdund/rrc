package ies

import "rrc/utils"

// BandParametersTxSL-r14 ::= SEQUENCE
type BandparameterstxslR14 struct {
	V2xBandwidthclasstxslR14 V2xBandwidthclassslR14
	V2xEnbScheduledR14       *utils.ENUMERATED
	V2xHighpowerR14          *utils.ENUMERATED
}
