package ies

import "rrc/utils"

// Inband-DifferentPCI-NB-r13 ::= SEQUENCE
type InbandDifferentpciNbR13 struct {
	EutraNumcrsPortsR13 InbandDifferentpciNbR13EutraNumcrsPortsR13
	RasteroffsetR13     ChannelrasteroffsetNbR13
	Spare               utils.BITSTRING `lb:2,ub:2`
}
