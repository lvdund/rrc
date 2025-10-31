package ies

import "rrc/utils"

// PUSCH-Config-pusch-RepTypeIndicatorDCI-0-1-r16 ::= ENUMERATED
type PuschConfigPuschReptypeindicatordci01R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigPuschReptypeindicatordci01R16EnumeratedNothing = iota
	PuschConfigPuschReptypeindicatordci01R16EnumeratedPusch_Reptypea
	PuschConfigPuschReptypeindicatordci01R16EnumeratedPusch_Reptypeb
)
