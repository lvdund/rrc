package ies

// SL-EUTRA-AnchorCarrierFreqList-r16 ::= SEQUENCE OF ARFCN-ValueEUTRA
// SIZE (1..maxFreqSL-EUTRA-r16)
type SlEutraAnchorcarrierfreqlistR16 struct {
	Value []ArfcnValueeutra `lb:1,ub:maxFreqSLEutraR16`
}
