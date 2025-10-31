package ies

import "rrc/utils"

// AccessStratumRelease ::= utils.ENUMERATED // Extensible
type Accessstratumrelease struct {
	Value utils.ENUMERATED
}

const (
	AccessstratumreleaseEnumeratedNothing = iota
	AccessstratumreleaseEnumeratedRel15
	AccessstratumreleaseEnumeratedRel16
	AccessstratumreleaseEnumeratedRel17
	AccessstratumreleaseEnumeratedSpare5
	AccessstratumreleaseEnumeratedSpare4
	AccessstratumreleaseEnumeratedSpare3
	AccessstratumreleaseEnumeratedSpare2
	AccessstratumreleaseEnumeratedSpare1
)
