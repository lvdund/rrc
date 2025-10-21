package ies

import "rrc/utils"

// CarrierFreqEUTRA-v9e0 ::= SEQUENCE
type CarrierfreqeutraV9e0 struct {
	DlCarrierfreqV9e0 ArfcnValueeutraR9
	UlCarrierfreqV9e0 *ArfcnValueeutraR9
}
