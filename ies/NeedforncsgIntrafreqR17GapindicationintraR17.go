package ies

import "rrc/utils"

// NeedForNCSG-IntraFreq-r17-gapIndicationIntra-r17 ::= ENUMERATED
type NeedforncsgIntrafreqR17GapindicationintraR17 struct {
	Value utils.ENUMERATED
}

const (
	NeedforncsgIntrafreqR17GapindicationintraR17EnumeratedNothing = iota
	NeedforncsgIntrafreqR17GapindicationintraR17EnumeratedGap
	NeedforncsgIntrafreqR17GapindicationintraR17EnumeratedNcsg
	NeedforncsgIntrafreqR17GapindicationintraR17EnumeratedNogap_Noncsg
)
