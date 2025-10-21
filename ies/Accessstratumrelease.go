package ies

import "rrc/utils"

// AccessStratumRelease ::= ENUMERATED
// Extensible
type Accessstratumrelease struct {
	Value utils.ENUMERATED
}

const (
	AccessstratumreleaseRel8  = 0
	AccessstratumreleaseRel9  = 1
	AccessstratumreleaseRel10 = 2
	AccessstratumreleaseRel11 = 3
	AccessstratumreleaseRel12 = 4
	AccessstratumreleaseRel13 = 5
	AccessstratumreleaseRel14 = 6
	AccessstratumreleaseRel15 = 7
	AccessstratumreleaseRel16 = 8
)
