package ies

// MCC ::= SEQUENCE OF MCC-MNC-Digit
// SIZE (3)
type Mcc struct {
	Value []MccMncDigit `lb:3,ub:3`
}
