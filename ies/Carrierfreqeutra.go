package ies

// CarrierFreqEUTRA ::= SEQUENCE
type Carrierfreqeutra struct {
	DlCarrierfreq ArfcnValueeutra
	UlCarrierfreq *ArfcnValueeutra
}
