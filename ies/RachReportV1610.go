package ies

import "rrc/utils"

// RACH-Report-v1610 ::= SEQUENCE
type RachReportV1610 struct {
	InitialcelR16  utils.INTEGER
	EdtFallbackR16 bool
}
