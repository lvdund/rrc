package ies

import "rrc/utils"

// AffectedCarrierFreqList-r11 ::= SEQUENCE OF AffectedCarrierFreq-r11
// SIZE (1..maxFreqIDC-r11)
type AffectedcarrierfreqlistR11 struct {
	Value []AffectedcarrierfreqR11
}
