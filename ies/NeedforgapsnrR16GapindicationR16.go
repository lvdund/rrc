package ies

import "rrc/utils"

// NeedForGapsNR-r16-gapIndication-r16 ::= ENUMERATED
type NeedforgapsnrR16GapindicationR16 struct {
	Value utils.ENUMERATED
}

const (
	NeedforgapsnrR16GapindicationR16EnumeratedNothing = iota
	NeedforgapsnrR16GapindicationR16EnumeratedGap
	NeedforgapsnrR16GapindicationR16EnumeratedNo_Gap
)
