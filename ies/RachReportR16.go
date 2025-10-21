package ies

import "rrc/utils"

// RACH-Report-r16 ::= SEQUENCE
type RachReportR16 struct {
	NumberofpreamblessentR16 NumberofpreamblessentR11
	ContentiondetectedR16    bool
}
