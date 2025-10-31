package ies

// AffectedCarrierFreqCombNR ::= SEQUENCE OF ARFCN-ValueNR
// SIZE (1..maxNrofServingCells)
type Affectedcarrierfreqcombnr struct {
	Value []ArfcnValuenr `lb:1,ub:maxNrofServingCells`
}
