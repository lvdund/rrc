package ies

import "rrc/utils"

// PerCC-GapIndication-r14-gapIndication-r14 ::= ENUMERATED
type PerccGapindicationR14GapindicationR14 struct {
	Value utils.ENUMERATED
}

const (
	PerccGapindicationR14GapindicationR14EnumeratedNothing = iota
	PerccGapindicationR14GapindicationR14EnumeratedGap
	PerccGapindicationR14GapindicationR14EnumeratedNcsg
	PerccGapindicationR14GapindicationR14EnumeratedNogap_Noncsg
)
