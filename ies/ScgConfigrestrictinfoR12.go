package ies

import "rrc/utils"

// SCG-ConfigRestrictInfo-r12 ::= SEQUENCE
type ScgConfigrestrictinfoR12 struct {
	MaxschTbBitsdlR12 utils.INTEGER
	MaxschTbBitsulR12 utils.INTEGER
}
