package ies

// SL-AnchorCarrierFreqList-V2X-r14 ::= SEQUENCE OF ARFCN-ValueEUTRA-r9
// SIZE (1..maxFreqV2X-r14)
type SlAnchorcarrierfreqlistV2xR14 struct {
	Value []ArfcnValueeutraR9 `lb:1,ub:maxFreqV2XR14`
}
