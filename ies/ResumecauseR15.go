package ies

import "rrc/utils"

// ResumeCause-r15 ::= ENUMERATED
type ResumecauseR15 struct {
	Value utils.ENUMERATED
}

const (
	ResumecauseR15Emergency          = 0
	ResumecauseR15Highpriorityaccess = 1
	ResumecauseR15MtAccess           = 2
	ResumecauseR15MoSignalling       = 3
	ResumecauseR15MoData             = 4
	ResumecauseR15RnaUpdate          = 5
	ResumecauseR15MoVoicecall        = 6
	ResumecauseR15Spare1             = 7
)
