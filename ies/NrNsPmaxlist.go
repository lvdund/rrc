package ies

// NR-NS-PmaxList ::= SEQUENCE OF NR-NS-PmaxValue
// SIZE (1..maxNR-NS-Pmax)
type NrNsPmaxlist struct {
	Value []NrNsPmaxvalue `lb:1,ub:maxNRNsPmax`
}
