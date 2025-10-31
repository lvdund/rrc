package ies

import "rrc/utils"

// AdditionalSpectrumEmissionNR-r15 ::= utils.INTEGER (0..7)
type AdditionalspectrumemissionnrR15 struct {
	Value utils.INTEGER `lb:0,ub:7`
}
