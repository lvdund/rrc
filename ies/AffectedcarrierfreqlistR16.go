package ies

// AffectedCarrierFreqList-r16 ::= SEQUENCE OF AffectedCarrierFreq-r16
// SIZE (1.. maxFreqIDC-r16)
type AffectedcarrierfreqlistR16 struct {
	Value []AffectedcarrierfreqR16 `lb:1,ub:maxFreqIDCR16`
}
