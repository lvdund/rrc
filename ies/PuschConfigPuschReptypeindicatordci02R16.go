package ies

import "rrc/utils"

// PUSCH-Config-pusch-RepTypeIndicatorDCI-0-2-r16 ::= ENUMERATED
type PuschConfigPuschReptypeindicatordci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigPuschReptypeindicatordci02R16EnumeratedNothing = iota
	PuschConfigPuschReptypeindicatordci02R16EnumeratedPusch_Reptypea
	PuschConfigPuschReptypeindicatordci02R16EnumeratedPusch_Reptypeb
)
