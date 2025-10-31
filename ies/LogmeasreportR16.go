package ies

import "rrc/utils"

// LogMeasReport-r16 ::= SEQUENCE
// Extensible
type LogmeasreportR16 struct {
	AbsolutetimestampR16        AbsolutetimeinfoR16
	TracereferenceR16           TracereferenceR16
	TracerecordingsessionrefR16 utils.OCTETSTRING `lb:2,ub:2`
	TceIdR16                    utils.OCTETSTRING `lb:1,ub:1`
	LogmeasinfolistR16          LogmeasinfolistR16
	LogmeasavailableR16         *utils.BOOLEAN
	LogmeasavailablebtR16       *utils.BOOLEAN
	LogmeasavailablewlanR16     *utils.BOOLEAN
}
