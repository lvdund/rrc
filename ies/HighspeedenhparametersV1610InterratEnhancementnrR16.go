package ies

import "rrc/utils"

// HighSpeedEnhParameters-v1610-interRAT-enhancementNR-r16 ::= ENUMERATED
type HighspeedenhparametersV1610InterratEnhancementnrR16 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedenhparametersV1610InterratEnhancementnrR16EnumeratedNothing = iota
	HighspeedenhparametersV1610InterratEnhancementnrR16EnumeratedSupported
)
