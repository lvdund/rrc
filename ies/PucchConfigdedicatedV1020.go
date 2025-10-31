package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1020 ::= SEQUENCE
type PucchConfigdedicatedV1020 struct {
	PucchFormatR10                            *PucchConfigdedicatedV1020PucchFormatR10
	TwoantennaportactivatedpucchFormat1a1bR10 *bool
	SimultaneouspucchPuschR10                 *bool
	N1pucchAnRepp1R10                         *utils.INTEGER `lb:0,ub:2047`
}
