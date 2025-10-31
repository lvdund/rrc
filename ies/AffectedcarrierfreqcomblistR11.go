package ies

// AffectedCarrierFreqCombList-r11 ::= SEQUENCE OF AffectedCarrierFreqComb-r11
// SIZE (1..maxCombIDC-r11)
type AffectedcarrierfreqcomblistR11 struct {
	Value []AffectedcarrierfreqcombR11 `lb:1,ub:maxCombIDCR11`
}
