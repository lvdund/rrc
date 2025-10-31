package ies

import "rrc/utils"

// RACH-Report-v1610 ::= SEQUENCE
type RachReportV1610 struct {
	InitialcelR16  utils.INTEGER `lb:0,ub:3`
	EdtFallbackR16 utils.BOOLEAN
}
