package ies

// AffectedCarrierFreqCombList-r16 ::= SEQUENCE OF AffectedCarrierFreqComb-r16
// SIZE (1..maxCombIDC-r16)
type AffectedcarrierfreqcomblistR16 struct {
	Value []AffectedcarrierfreqcombR16 `lb:1,ub:maxCombIDCR16`
}
