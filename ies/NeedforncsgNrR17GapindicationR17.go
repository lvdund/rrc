package ies

import "rrc/utils"

// NeedForNCSG-NR-r17-gapIndication-r17 ::= ENUMERATED
type NeedforncsgNrR17GapindicationR17 struct {
	Value utils.ENUMERATED
}

const (
	NeedforncsgNrR17GapindicationR17EnumeratedNothing = iota
	NeedforncsgNrR17GapindicationR17EnumeratedGap
	NeedforncsgNrR17GapindicationR17EnumeratedNcsg
	NeedforncsgNrR17GapindicationR17EnumeratedNogap_Noncsg
)
