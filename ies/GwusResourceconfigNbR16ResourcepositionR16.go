package ies

import "rrc/utils"

// GWUS-ResourceConfig-NB-r16-resourcePosition-r16 ::= ENUMERATED
type GwusResourceconfigNbR16ResourcepositionR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusResourceconfigNbR16ResourcepositionR16EnumeratedNothing = iota
	GwusResourceconfigNbR16ResourcepositionR16EnumeratedPrimary
	GwusResourceconfigNbR16ResourcepositionR16EnumeratedSecondary
)
