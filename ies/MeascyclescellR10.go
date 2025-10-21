package ies

import "rrc/utils"

// MeasCycleSCell-r10 ::= ENUMERATED
type MeascyclescellR10 struct {
	Value utils.ENUMERATED
}

const (
	MeascyclescellR10Sf160  = 0
	MeascyclescellR10Sf256  = 1
	MeascyclescellR10Sf320  = 2
	MeascyclescellR10Sf512  = 3
	MeascyclescellR10Sf640  = 4
	MeascyclescellR10Sf1024 = 5
	MeascyclescellR10Sf1280 = 6
	MeascyclescellR10Spare1 = 7
)
