package ies

import "rrc/utils"

// WUS-Config-r15-maxDurationFactor-r15 ::= ENUMERATED
type WusConfigR15MaxdurationfactorR15 struct {
	Value utils.ENUMERATED
}

const (
	WusConfigR15MaxdurationfactorR15EnumeratedNothing = iota
	WusConfigR15MaxdurationfactorR15EnumeratedOne32th
	WusConfigR15MaxdurationfactorR15EnumeratedOne16th
	WusConfigR15MaxdurationfactorR15EnumeratedOne8th
	WusConfigR15MaxdurationfactorR15EnumeratedOne4th
)
