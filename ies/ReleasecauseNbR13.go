package ies

import "rrc/utils"

// ReleaseCause-NB-r13 ::= ENUMERATED
type ReleasecauseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	ReleasecauseNbR13Loadbalancingtaurequired = 0
	ReleasecauseNbR13Other                    = 1
	ReleasecauseNbR13RrcSuspend               = 2
	ReleasecauseNbR13Spare1                   = 3
)
