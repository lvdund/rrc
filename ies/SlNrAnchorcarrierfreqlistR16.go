package ies

// SL-NR-AnchorCarrierFreqList-r16 ::= SEQUENCE OF ARFCN-ValueNR
// SIZE (1..maxFreqSL-NR-r16)
type SlNrAnchorcarrierfreqlistR16 struct {
	Value []ArfcnValuenr `lb:1,ub:maxFreqSLNrR16`
}
