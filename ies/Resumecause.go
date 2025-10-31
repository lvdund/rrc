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
	ResumecauseEnumeratedMo_Voicecall
	ResumecauseEnumeratedMo_Videocall
	ResumecauseEnumeratedMo_Sms
	ResumecauseEnumeratedRna_Update
	ResumecauseEnumeratedMps_Priorityaccess
	ResumecauseEnumeratedMcs_Priorityaccess
	ResumecauseEnumeratedSpare1
	ResumecauseEnumeratedSpare2
	ResumecauseEnumeratedSpare3
	ResumecauseEnumeratedSpare4
	ResumecauseEnumeratedSpare5
)
