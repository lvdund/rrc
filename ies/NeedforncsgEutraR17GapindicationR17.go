package ies

import "rrc/utils"

// NeedForNCSG-EUTRA-r17-gapIndication-r17 ::= ENUMERATED
type NeedforncsgEutraR17GapindicationR17 struct {
	Value utils.ENUMERATED
}

const (
	NeedforncsgEutraR17GapindicationR17EnumeratedNothing = iota
	NeedforncsgEutraR17GapindicationR17EnumeratedGap
	NeedforncsgEutraR17GapindicationR17EnumeratedNcsg
	NeedforncsgEutraR17GapindicationR17EnumeratedNogap_Noncsg
)
