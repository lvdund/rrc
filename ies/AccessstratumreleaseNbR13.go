package ies

import "rrc/utils"

// AccessStratumRelease-NB-r13 ::= ENUMERATED
// Extensible
type AccessstratumreleaseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	AccessstratumreleaseNbR13Rel13  = 0
	AccessstratumreleaseNbR13Rel14  = 1
	AccessstratumreleaseNbR13Rel15  = 2
	AccessstratumreleaseNbR13Rel16  = 3
	AccessstratumreleaseNbR13Spare4 = 4
	AccessstratumreleaseNbR13Spare3 = 5
	AccessstratumreleaseNbR13Spare2 = 6
	AccessstratumreleaseNbR13Spare1 = 7
)
