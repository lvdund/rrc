package ies

import "rrc/utils"

// CarrierFreqsGERAN-NB-r16-peo-r16 ::= ENUMERATED
type CarrierfreqsgeranNbR16PeoR16 struct {
	Value utils.ENUMERATED
}

const (
	CarrierfreqsgeranNbR16PeoR16EnumeratedNothing = iota
	CarrierfreqsgeranNbR16PeoR16EnumeratedSupported
)
