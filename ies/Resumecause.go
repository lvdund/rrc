package ies

import "rrc/utils"

// ResumeCause ::= ENUMERATED
type Resumecause struct {
	Value utils.ENUMERATED
}

const (
	ResumecauseEnumeratedNothing = iota
	ResumecauseEnumeratedEmergency
	ResumecauseEnumeratedHighpriorityaccess
	ResumecauseEnumeratedMt_Access
	ResumecauseEnumeratedMo_Signalling
	ResumecauseEnumeratedMo_Data
	ResumecauseEnumeratedDelaytolerantaccess_V1020
	ResumecauseEnumeratedMo_Voicecall_V1280
	ResumecauseEnumeratedMt_Edt_V1610
)
