package ies

import "rrc/utils"

// Inband-DifferentPCI-NB-r13 ::= SEQUENCE
type InbandDifferentpciNbR13 struct {
	EutraNumcrsPortsR13 utils.ENUMERATED
	RasteroffsetR13     ChannelrasteroffsetNbR13
	Spare               utils.BITSTRING
}
