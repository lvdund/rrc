package ies

import "rrc/utils"

// CarrierFreqEUTRA-NB-r16-sib1-r16 ::= ENUMERATED
type CarrierfreqeutraNbR16Sib1R16 struct {
	Value utils.ENUMERATED
}

const (
	CarrierfreqeutraNbR16Sib1R16EnumeratedNothing = iota
	CarrierfreqeutraNbR16Sib1R16EnumeratedSupported
)
