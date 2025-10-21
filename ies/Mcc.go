package ies

import "rrc/utils"

// MCC ::= SEQUENCE OF MCC-MNC-Digit
// SIZE (3)
type Mcc struct {
	Value []MccMncDigit
}
