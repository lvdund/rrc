package ies

import "rrc/utils"

// FreqPrioritySlicing-r17 ::= SEQUENCE
type FreqpriorityslicingR17 struct {
	DlImplicitcarrierfreqR17 utils.INTEGER `lb:0,ub:maxFreq`
	SliceinfolistR17         *SliceinfolistR17
}
