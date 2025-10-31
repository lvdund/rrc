package ies

import "rrc/utils"

// SCG-ConfigRestrictInfo-r12 ::= SEQUENCE
type ScgConfigrestrictinfoR12 struct {
	MaxschTbBitsdlR12 utils.INTEGER `lb:0,ub:100`
	MaxschTbBitsulR12 utils.INTEGER `lb:0,ub:100`
}
