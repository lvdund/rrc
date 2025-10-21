package ies

import "rrc/utils"

// MeasSubframePattern-r10 ::= CHOICE
// Extensible
type MeassubframepatternR10 interface {
	isMeassubframepatternR10()
}

type MeassubframepatternR10SubframepatternfddR10 struct {
	Value utils.BITSTRING
}

func (*MeassubframepatternR10SubframepatternfddR10) isMeassubframepatternR10() {}

type MeassubframepatternR10SubframepatterntddR10 struct {
	Value interface{}
}

func (*MeassubframepatternR10SubframepatterntddR10) isMeassubframepatternR10() {}
