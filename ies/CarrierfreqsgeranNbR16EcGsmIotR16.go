package ies

import "rrc/utils"

// CarrierFreqsGERAN-NB-r16-ec-GSM-IOT-r16 ::= ENUMERATED
type CarrierfreqsgeranNbR16EcGsmIotR16 struct {
	Value utils.ENUMERATED
}

const (
	CarrierfreqsgeranNbR16EcGsmIotR16EnumeratedNothing = iota
	CarrierfreqsgeranNbR16EcGsmIotR16EnumeratedSupported
)
