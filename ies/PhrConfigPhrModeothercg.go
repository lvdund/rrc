package ies

import "rrc/utils"

// PHR-Config-phr-ModeOtherCG ::= ENUMERATED
type PhrConfigPhrModeothercg struct {
	Value utils.ENUMERATED
}

const (
	PhrConfigPhrModeothercgEnumeratedNothing = iota
	PhrConfigPhrModeothercgEnumeratedReal
	PhrConfigPhrModeothercgEnumeratedVirtual
)
