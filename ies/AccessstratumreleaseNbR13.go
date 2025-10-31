package ies

import "rrc/utils"

// AccessStratumRelease-NB-r13 ::= utils.ENUMERATED // Extensible
type AccessstratumreleaseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	AccessstratumreleaseNbR13EnumeratedNothing = iota
	AccessstratumreleaseNbR13EnumeratedRel13
	AccessstratumreleaseNbR13EnumeratedRel14
	AccessstratumreleaseNbR13EnumeratedRel15
	AccessstratumreleaseNbR13EnumeratedRel16
	AccessstratumreleaseNbR13EnumeratedSpare4
	AccessstratumreleaseNbR13EnumeratedSpare3
	AccessstratumreleaseNbR13EnumeratedSpare2
	AccessstratumreleaseNbR13EnumeratedSpare1
)
