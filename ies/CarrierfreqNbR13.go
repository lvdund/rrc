package ies

import "rrc/utils"

// CarrierFreq-NB-r13 ::= SEQUENCE
type CarrierfreqNbR13 struct {
	CarrierfreqR13       ArfcnValueeutraR9
	CarrierfreqoffsetR13 *utils.ENUMERATED
}
