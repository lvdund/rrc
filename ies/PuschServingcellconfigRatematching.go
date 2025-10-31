package ies

import "rrc/utils"

// PUSCH-ServingCellConfig-rateMatching ::= ENUMERATED
type PuschServingcellconfigRatematching struct {
	Value utils.ENUMERATED
}

const (
	PuschServingcellconfigRatematchingEnumeratedNothing = iota
	PuschServingcellconfigRatematchingEnumeratedLimitedbufferrm
)
