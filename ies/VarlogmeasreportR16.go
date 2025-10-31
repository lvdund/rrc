package ies

import "rrc/utils"

// VarLogMeasReport-r16 ::= SEQUENCE
type VarlogmeasreportR16 struct {
	AbsolutetimeinfoR16         AbsolutetimeinfoR16
	TracereferenceR16           TracereferenceR16
	TracerecordingsessionrefR16 utils.OCTETSTRING `lb:2,ub:2`
	TceIdR16                    utils.OCTETSTRING `lb:1,ub:1`
	LogmeasinfolistR16          LogmeasinfolistR16
	PlmnIdentitylistR16         PlmnIdentitylist2R16
	SigloggedmeastypeR17        utils.BOOLEAN
}
