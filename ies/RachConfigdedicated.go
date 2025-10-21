package ies

import "rrc/utils"

// RACH-ConfigDedicated ::= SEQUENCE
type RachConfigdedicated struct {
	RaPreambleindex  utils.INTEGER
	RaPrachMaskindex utils.INTEGER
}
