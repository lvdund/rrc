package ies

import "rrc/utils"

// CarrierFreqEUTRA-NB-r16-sib1-BR-r16 ::= ENUMERATED
type CarrierfreqeutraNbR16Sib1BrR16 struct {
	Value utils.ENUMERATED
}

const (
	CarrierfreqeutraNbR16Sib1BrR16EnumeratedNothing = iota
	CarrierfreqeutraNbR16Sib1BrR16EnumeratedSupported
)
