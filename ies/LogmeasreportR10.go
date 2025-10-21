package ies

import "rrc/utils"

// LogMeasReport-r10 ::= SEQUENCE
// Extensible
type LogmeasreportR10 struct {
	AbsolutetimestampR10        AbsolutetimeinfoR10
	TracereferenceR10           TracereferenceR10
	TracerecordingsessionrefR10 utils.OCTETSTRING
	TceIdR10                    utils.OCTETSTRING
	LogmeasinfolistR10          LogmeasinfolistR10
	LogmeasavailableR10         *utils.ENUMERATED
}
