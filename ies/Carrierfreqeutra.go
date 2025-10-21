package ies

import "rrc/utils"

// CarrierFreqEUTRA ::= SEQUENCE
type Carrierfreqeutra struct {
	DlCarrierfreq ArfcnValueeutra
	UlCarrierfreq *ArfcnValueeutra
}
