package ies

import "rrc/utils"

// ReleaseCause ::= ENUMERATED
type Releasecause struct {
	Value utils.ENUMERATED
}

const (
	ReleasecauseLoadbalancingtaurequired    = 0
	ReleasecauseOther                       = 1
	ReleasecauseCsFallbackhighpriorityV1020 = 2
	ReleasecauseRrcSuspendV1320             = 3
)
