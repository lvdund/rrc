package ies

import "rrc/utils"

// ReleaseCause-NB-r13 ::= ENUMERATED
type ReleasecauseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	ReleasecauseNbR13EnumeratedNothing = iota
	ReleasecauseNbR13EnumeratedLoadbalancingtaurequired
	ReleasecauseNbR13EnumeratedOther
	ReleasecauseNbR13EnumeratedRrc_Suspend
	ReleasecauseNbR13EnumeratedSpare1
)
