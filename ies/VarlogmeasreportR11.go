package ies

import "rrc/utils"

// VarLogMeasReport-r11 ::= SEQUENCE
type VarlogmeasreportR11 struct {
	TracereferenceR10           TracereferenceR10
	TracerecordingsessionrefR10 utils.OCTETSTRING
	TceIdR10                    utils.OCTETSTRING
	PlmnIdentitylistR11         PlmnIdentitylist3R11
	AbsolutetimeinfoR10         AbsolutetimeinfoR10
	LogmeasinfolistR10          Logmeasinfolist2R10
}
