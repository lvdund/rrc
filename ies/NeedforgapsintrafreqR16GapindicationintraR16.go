package ies

import "rrc/utils"

// NeedForGapsIntraFreq-r16-gapIndicationIntra-r16 ::= ENUMERATED
type NeedforgapsintrafreqR16GapindicationintraR16 struct {
	Value utils.ENUMERATED
}

const (
	NeedforgapsintrafreqR16GapindicationintraR16EnumeratedNothing = iota
	NeedforgapsintrafreqR16GapindicationintraR16EnumeratedGap
	NeedforgapsintrafreqR16GapindicationintraR16EnumeratedNo_Gap
)
