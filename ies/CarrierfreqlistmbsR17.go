package ies

// CarrierFreqListMBS-r17 ::= SEQUENCE OF ARFCN-ValueNR
// SIZE (1..maxFreqMBS-r17)
type CarrierfreqlistmbsR17 struct {
	Value []ArfcnValuenr `lb:1,ub:maxFreqMBSR17`
}
