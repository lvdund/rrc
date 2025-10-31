package ies

import "rrc/utils"

// VarLogMeasReport-r10 ::= SEQUENCE
type VarlogmeasreportR10 struct {
	TracereferenceR10           TracereferenceR10
	TracerecordingsessionrefR10 utils.OCTETSTRING `lb:2,ub:2`
	TceIdR10                    utils.OCTETSTRING `lb:1,ub:1`
	PlmnIdentityR10             PlmnIdentity
	AbsolutetimeinfoR10         AbsolutetimeinfoR10
	LogmeasinfolistR10          Logmeasinfolist2R10
}
