package ies

import "rrc/utils"

// Guardband-NB-r13 ::= SEQUENCE
type GuardbandNbR13 struct {
	RasteroffsetR13 ChannelrasteroffsetNbR13
	Spare           utils.BITSTRING
}
