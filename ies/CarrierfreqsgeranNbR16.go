package ies

// CarrierFreqsGERAN-NB-r16 ::= SEQUENCE
// Extensible
type CarrierfreqsgeranNbR16 struct {
	CarrierfreqsR16 Carrierfreqsgeran
	EcGsmIotR16     *CarrierfreqsgeranNbR16EcGsmIotR16
	PeoR16          *CarrierfreqsgeranNbR16PeoR16
}
