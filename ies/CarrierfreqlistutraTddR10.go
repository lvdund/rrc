package ies

// CarrierFreqListUTRA-TDD-r10 ::= SEQUENCE OF ARFCN-ValueUTRA
// SIZE (1..maxFreqUTRA-TDD-r10)
type CarrierfreqlistutraTddR10 struct {
	Value []ArfcnValueutra `lb:1,ub:maxFreqUTRATddR10`
}
