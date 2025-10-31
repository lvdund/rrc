package ies

// AffectedCarrierFreqCombEUTRA ::= SEQUENCE OF ARFCN-ValueEUTRA
// SIZE (1..maxNrofServingCellsEUTRA)
type Affectedcarrierfreqcombeutra struct {
	Value []ArfcnValueeutra `lb:1,ub:maxNrofServingCellsEUTRA`
}
