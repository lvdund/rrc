package ies

import "rrc/utils"

// MeasSubframePattern-r10-subframePatternTDD-r10 ::= CHOICE
// Extensible
const (
	MeassubframepatternR10SubframepatterntddR10ChoiceNothing = iota
	MeassubframepatternR10SubframepatterntddR10ChoiceSubframeconfig15R10
	MeassubframepatternR10SubframepatterntddR10ChoiceSubframeconfig0R10
	MeassubframepatternR10SubframepatterntddR10ChoiceSubframeconfig6R10
)

type MeassubframepatternR10SubframepatterntddR10 struct {
	Choice              uint64
	Subframeconfig15R10 *utils.BITSTRING `lb:20,ub:20`
	Subframeconfig0R10  *utils.BITSTRING `lb:70,ub:70`
	Subframeconfig6R10  *utils.BITSTRING `lb:60,ub:60`
}
