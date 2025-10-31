package ies

// CarrierFreqEUTRA-NB-r16 ::= SEQUENCE
// Extensible
type CarrierfreqeutraNbR16 struct {
	CarrierfreqR16 ArfcnValueeutraR9
	Sib1R16        *CarrierfreqeutraNbR16Sib1R16
	Sib1BrR16      *CarrierfreqeutraNbR16Sib1BrR16
}
