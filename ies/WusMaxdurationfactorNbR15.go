package ies

import "rrc/utils"

// WUS-MaxDurationFactor-NB-r15 ::= ENUMERATED
type WusMaxdurationfactorNbR15 struct {
	Value utils.ENUMERATED
}

const (
	WusMaxdurationfactorNbR15EnumeratedNothing = iota
	WusMaxdurationfactorNbR15EnumeratedOne128th
	WusMaxdurationfactorNbR15EnumeratedOne64th
	WusMaxdurationfactorNbR15EnumeratedOne32th
	WusMaxdurationfactorNbR15EnumeratedOne16th
	WusMaxdurationfactorNbR15EnumeratedOneeighth
	WusMaxdurationfactorNbR15EnumeratedOnequarter
	WusMaxdurationfactorNbR15EnumeratedOnehalf
)
