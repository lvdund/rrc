package ies

// AffectedCarrierFreq-r11 ::= SEQUENCE
type AffectedcarrierfreqR11 struct {
	CarrierfreqR11           Measobjectid
	InterferencedirectionR11 AffectedcarrierfreqR11InterferencedirectionR11
}
