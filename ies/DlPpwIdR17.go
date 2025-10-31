package ies

import "rrc/utils"

// DL-PPW-ID-r17 ::= utils.INTEGER (0..maxNrofPPW-ID-1-r17)
type DlPpwIdR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPPWId1R17`
}
