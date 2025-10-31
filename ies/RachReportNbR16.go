package ies

import "rrc/utils"

// RACH-Report-NB-r16 ::= SEQUENCE
type RachReportNbR16 struct {
	NumberofpreamblessentR16 utils.INTEGER `lb:0,ub:64`
	ContentiondetectedR16    utils.BOOLEAN
	InitialnrsrpLevelR16     utils.INTEGER `lb:0,ub:2`
	EdtFallbackR16           utils.BOOLEAN
}
