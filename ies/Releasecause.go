package ies

import "rrc/utils"

// ReleaseCause ::= ENUMERATED
type Releasecause struct {
	Value utils.ENUMERATED
}

const (
	ReleasecauseEnumeratedNothing = iota
	ReleasecauseEnumeratedLoadbalancingtaurequired
	ReleasecauseEnumeratedOther
	ReleasecauseEnumeratedCs_Fallbackhighpriority_V1020
	ReleasecauseEnumeratedRrc_Suspend_V1320
)
