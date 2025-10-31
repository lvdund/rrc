package ies

// FreqPriorityListUTRA-TDD ::= SEQUENCE OF FreqPriorityUTRA-TDD
// SIZE (1..maxUTRA-TDD-Carrier)
type FreqprioritylistutraTdd struct {
	Value []FreqpriorityutraTdd `lb:1,ub:maxUTRATddCarrier`
}
