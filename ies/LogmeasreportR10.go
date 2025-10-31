package ies

import "rrc/utils"

// LogMeasReport-r10 ::= SEQUENCE
// Extensible
type LogmeasreportR10 struct {
	AbsolutetimestampR10        AbsolutetimeinfoR10
	TracereferenceR10           TracereferenceR10
	TracerecordingsessionrefR10 utils.OCTETSTRING `lb:2,ub:2`
	TceIdR10                    utils.OCTETSTRING `lb:1,ub:1`
	LogmeasinfolistR10          LogmeasinfolistR10
	LogmeasavailableR10         *bool
}
