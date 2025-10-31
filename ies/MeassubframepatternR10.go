package ies

import "rrc/utils"

// MeasSubframePattern-r10 ::= CHOICE
// Extensible
const (
	MeassubframepatternR10ChoiceNothing = iota
	MeassubframepatternR10ChoiceSubframepatternfddR10
	MeassubframepatternR10ChoiceSubframepatterntddR10
)

type MeassubframepatternR10 struct {
	Choice                uint64
	SubframepatternfddR10 *utils.BITSTRING `lb:40,ub:40`
	SubframepatterntddR10 *MeassubframepatternR10SubframepatterntddR10
}
