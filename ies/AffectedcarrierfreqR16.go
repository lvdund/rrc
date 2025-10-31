package ies

// AffectedCarrierFreq-r16 ::= SEQUENCE
type AffectedcarrierfreqR16 struct {
	CarrierfreqR16           ArfcnValuenr
	InterferencedirectionR16 AffectedcarrierfreqR16InterferencedirectionR16
}
