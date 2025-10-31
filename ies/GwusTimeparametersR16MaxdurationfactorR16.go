package ies

import "rrc/utils"

// GWUS-TimeParameters-r16-maxDurationFactor-r16 ::= ENUMERATED
type GwusTimeparametersR16MaxdurationfactorR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusTimeparametersR16MaxdurationfactorR16EnumeratedNothing = iota
	GwusTimeparametersR16MaxdurationfactorR16EnumeratedOne32th
	GwusTimeparametersR16MaxdurationfactorR16EnumeratedOne16th
	GwusTimeparametersR16MaxdurationfactorR16EnumeratedOne8th
	GwusTimeparametersR16MaxdurationfactorR16EnumeratedOne4th
)
