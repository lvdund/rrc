package ies

import "rrc/utils"

// AdditionalSpectrumEmission-v10l0 ::= utils.INTEGER (33..288)
type AdditionalspectrumemissionV10l0 struct {
	Value utils.INTEGER `lb:0,ub:288`
}
