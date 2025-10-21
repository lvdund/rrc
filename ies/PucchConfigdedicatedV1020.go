package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1020 ::= SEQUENCE
type PucchConfigdedicatedV1020 struct {
	PucchFormatR10                            *interface{}
	TwoantennaportactivatedpucchFormat1a1bR10 *utils.ENUMERATED
	SimultaneouspucchPuschR10                 *utils.ENUMERATED
	N1pucchAnRepp1R10                         *utils.INTEGER
}
