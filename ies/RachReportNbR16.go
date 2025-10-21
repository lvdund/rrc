package ies

import "rrc/utils"

// RACH-Report-NB-r16 ::= SEQUENCE
type RachReportNbR16 struct {
	NumberofpreamblessentR16 utils.INTEGER
	ContentiondetectedR16    bool
	InitialnrsrpLevelR16     utils.INTEGER
	EdtFallbackR16           bool
}
