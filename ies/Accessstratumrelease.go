package ies

import "rrc/utils"

// AccessStratumRelease ::= utils.ENUMERATED // Extensible
type Accessstratumrelease struct {
	Value utils.ENUMERATED
}

const (
	AccessstratumreleaseEnumeratedNothing = iota
	AccessstratumreleaseEnumeratedRel8
	AccessstratumreleaseEnumeratedRel9
	AccessstratumreleaseEnumeratedRel10
	AccessstratumreleaseEnumeratedRel11
	AccessstratumreleaseEnumeratedRel12
	AccessstratumreleaseEnumeratedRel13
	AccessstratumreleaseEnumeratedRel14
	AccessstratumreleaseEnumeratedRel15
	AccessstratumreleaseEnumeratedRel16
)
