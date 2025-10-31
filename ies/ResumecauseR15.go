package ies

import "rrc/utils"

// ResumeCause-r15 ::= ENUMERATED
type ResumecauseR15 struct {
	Value utils.ENUMERATED
}

const (
	ResumecauseR15EnumeratedNothing = iota
	ResumecauseR15EnumeratedEmergency
	ResumecauseR15EnumeratedHighpriorityaccess
	ResumecauseR15EnumeratedMt_Access
	ResumecauseR15EnumeratedMo_Signalling
	ResumecauseR15EnumeratedMo_Data
	ResumecauseR15EnumeratedRna_Update
	ResumecauseR15EnumeratedMo_Voicecall
	ResumecauseR15EnumeratedSpare1
)
