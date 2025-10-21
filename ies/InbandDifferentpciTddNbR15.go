package ies

import "rrc/utils"

// Inband-DifferentPCI-TDD-NB-r15 ::= SEQUENCE
type InbandDifferentpciTddNbR15 struct {
	EutraNumcrsPortsR15  utils.ENUMERATED
	RasteroffsetR15      ChannelrasteroffsetNbR13
	SibInbandlocationR15 utils.ENUMERATED
	Spare                utils.BITSTRING
}
