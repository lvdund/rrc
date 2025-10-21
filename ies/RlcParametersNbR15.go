package ies

import "rrc/utils"

// RLC-Parameters-NB-r15 ::= SEQUENCE
type RlcParametersNbR15 struct {
	RlcUmR15 *utils.ENUMERATED
}
