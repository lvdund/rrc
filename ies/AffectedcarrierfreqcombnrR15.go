package ies

// AffectedCarrierFreqCombNR-r15 ::= SEQUENCE OF ARFCN-ValueNR-r15
// SIZE (1..maxServCellNR-r15)
type AffectedcarrierfreqcombnrR15 struct {
	Value []ArfcnValuenrR15 `lb:1,ub:maxServCellNRR15`
}
