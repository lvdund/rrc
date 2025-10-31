package ies

// AffectedCarrierFreqList-r11 ::= SEQUENCE OF AffectedCarrierFreq-r11
// SIZE (1..maxFreqIDC-r11)
type AffectedcarrierfreqlistR11 struct {
	Value []AffectedcarrierfreqR11 `lb:1,ub:maxFreqIDCR11`
}
