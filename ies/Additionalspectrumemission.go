package ies

import "rrc/utils"

// AdditionalSpectrumEmission ::= utils.INTEGER (1..32)
type Additionalspectrumemission struct {
	Value utils.INTEGER `lb:0,ub:32`
}
