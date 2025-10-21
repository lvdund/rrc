package ies

import "rrc/utils"

// CarrierFreqsGERAN-NB-r16 ::= SEQUENCE
// Extensible
type CarrierfreqsgeranNbR16 struct {
	CarrierfreqsR16 Carrierfreqsgeran
	EcGsmIotR16     *utils.ENUMERATED
	PeoR16          *utils.ENUMERATED
}
