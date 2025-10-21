package ies

import "rrc/utils"

// VarMeasReportList ::= SEQUENCE OF VarMeasReport
// SIZE (1..maxMeasId)
type Varmeasreportlist struct {
	Value []Varmeasreport
}
