package ies

import "rrc/utils"

// TCI-UL-State-Id-r17 ::= utils.INTEGER (0..maxUL-TCI-1-r17)
type TciUlStateIdR17 struct {
	Value utils.INTEGER `lb:0,ub:maxULTci1R17`
}
