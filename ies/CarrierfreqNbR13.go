package ies

// CarrierFreq-NB-r13 ::= SEQUENCE
type CarrierfreqNbR13 struct {
	CarrierfreqR13       ArfcnValueeutraR9
	CarrierfreqoffsetR13 *CarrierfreqNbR13CarrierfreqoffsetR13
}
