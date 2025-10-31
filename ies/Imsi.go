package ies

// IMSI ::= SEQUENCE OF IMSI-Digit
// SIZE (6..21)
type Imsi struct {
	Value []ImsiDigit `lb:6,ub:21`
}
