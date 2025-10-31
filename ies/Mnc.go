package ies

// MNC ::= SEQUENCE OF MCC-MNC-Digit
// SIZE (2..3)
type Mnc struct {
	Value []MccMncDigit `lb:2,ub:3`
}
