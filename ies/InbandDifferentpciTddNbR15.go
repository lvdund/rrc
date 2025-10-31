package ies

import "rrc/utils"

// Inband-DifferentPCI-TDD-NB-r15 ::= SEQUENCE
type InbandDifferentpciTddNbR15 struct {
	EutraNumcrsPortsR15  InbandDifferentpciTddNbR15EutraNumcrsPortsR15
	RasteroffsetR15      ChannelrasteroffsetNbR13
	SibInbandlocationR15 InbandDifferentpciTddNbR15SibInbandlocationR15
	Spare                utils.BITSTRING `lb:2,ub:2`
}
