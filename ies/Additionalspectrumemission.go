package ies

import "rrc/utils"

// AdditionalSpectrumEmission ::= utils.INTEGER (0..7)
type Additionalspectrumemission struct {
	Value utils.INTEGER `lb:0,ub:7`
}
