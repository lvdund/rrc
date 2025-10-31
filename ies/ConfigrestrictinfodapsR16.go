package ies

import "rrc/utils"

// ConfigRestrictInfoDAPS-r16 ::= SEQUENCE
type ConfigrestrictinfodapsR16 struct {
	MaxschTbBitsdlR16 *utils.INTEGER `lb:0,ub:100`
	MaxschTbBitsulR16 *utils.INTEGER `lb:0,ub:100`
}
