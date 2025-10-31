package ies

import "rrc/utils"

// VarLogMeasReport-r11 ::= SEQUENCE
type VarlogmeasreportR11 struct {
	TracereferenceR10           TracereferenceR10
	TracerecordingsessionrefR10 utils.OCTETSTRING `lb:2,ub:2`
	TceIdR10                    utils.OCTETSTRING `lb:1,ub:1`
	PlmnIdentitylistR11         PlmnIdentitylist3R11
	AbsolutetimeinfoR10         AbsolutetimeinfoR10
	LogmeasinfolistR10          Logmeasinfolist2R10
}
