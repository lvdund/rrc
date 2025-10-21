package ies

import "rrc/utils"

// WUS-MaxDurationFactor-NB-r15 ::= ENUMERATED
type WusMaxdurationfactorNbR15 struct {
	Value utils.ENUMERATED
}

const (
	WusMaxdurationfactorNbR15One128th   = 0
	WusMaxdurationfactorNbR15One64th    = 1
	WusMaxdurationfactorNbR15One32th    = 2
	WusMaxdurationfactorNbR15One16th    = 3
	WusMaxdurationfactorNbR15Oneeighth  = 4
	WusMaxdurationfactorNbR15Onequarter = 5
	WusMaxdurationfactorNbR15Onehalf    = 6
)
