package ies

import "rrc/utils"

// PUCCH-Format3-Conf-r13 ::= SEQUENCE
type PucchFormat3ConfR13 struct {
	N3pucchAnListR13                       *[]utils.INTEGER `lb:1,ub:4`
	TwoantennaportactivatedpucchFormat3R13 *PucchFormat3ConfR13TwoantennaportactivatedpucchFormat3R13
}
