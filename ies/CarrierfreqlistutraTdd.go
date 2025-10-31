package ies

// CarrierFreqListUTRA-TDD ::= SEQUENCE OF CarrierFreqUTRA-TDD
// SIZE (1..maxUTRA-TDD-Carrier)
type CarrierfreqlistutraTdd struct {
	Value []CarrierfrequtraTdd `lb:1,ub:maxUTRATddCarrier`
}
