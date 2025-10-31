package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-pucch-Format-r13-format3-r13 ::= SEQUENCE
type PucchConfigdedicatedR13PucchFormatR13Format3R13 struct {
	N3pucchAnListR13                       *[]utils.INTEGER `lb:1,ub:4`
	TwoantennaportactivatedpucchFormat3R13 *PucchConfigdedicatedR13PucchFormatR13Format3R13TwoantennaportactivatedpucchFormat3R13
}
