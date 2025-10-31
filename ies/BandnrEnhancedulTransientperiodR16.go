package ies

import "rrc/utils"

// BandNR-enhancedUL-TransientPeriod-r16 ::= ENUMERATED
type BandnrEnhancedulTransientperiodR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrEnhancedulTransientperiodR16EnumeratedNothing = iota
	BandnrEnhancedulTransientperiodR16EnumeratedUs2
	BandnrEnhancedulTransientperiodR16EnumeratedUs4
	BandnrEnhancedulTransientperiodR16EnumeratedUs7
)
