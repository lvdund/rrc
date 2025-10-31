package ies

// AffectedCarrierFreqCombList-r13 ::= SEQUENCE OF AffectedCarrierFreqComb-r13
// SIZE (1..maxCombIDC-r11)
type AffectedcarrierfreqcomblistR13 struct {
	Value []AffectedcarrierfreqcombR13 `lb:1,ub:maxCombIDCR11`
}
