package ies

// EUTRA-NS-PmaxList ::= SEQUENCE OF EUTRA-NS-PmaxValue
// SIZE (1..maxEUTRA-NS-Pmax)
type EutraNsPmaxlist struct {
	Value []EutraNsPmaxvalue `lb:1,ub:maxEUTRANsPmax`
}
