package ies

// CarrierFreqListUTRA-TDD-Ext-r12 ::= SEQUENCE OF CarrierFreqUTRA-TDD-r12
// SIZE (1..maxUTRA-TDD-Carrier)
type CarrierfreqlistutraTddExtR12 struct {
	Value []CarrierfrequtraTddR12 `lb:1,ub:maxUTRATddCarrier`
}
