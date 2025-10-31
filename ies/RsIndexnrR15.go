package ies

import "rrc/utils"

// RS-IndexNR-r15 ::= utils.INTEGER (0.. maxRS-Index-1-r15)
type RsIndexnrR15 struct {
	Value utils.INTEGER `lb:0,ub:maxRSIndex1R15`
}
