package ies

import "rrc/utils"

// RACH-ConfigDedicated ::= SEQUENCE
type RachConfigdedicated struct {
	RaPreambleindex  utils.INTEGER `lb:0,ub:63`
	RaPrachMaskindex utils.INTEGER `lb:0,ub:15`
}
