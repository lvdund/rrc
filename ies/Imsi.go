package ies

import "rrc/utils"

// IMSI ::= SEQUENCE OF IMSI-Digit
// SIZE (6..21)
type Imsi struct {
	Value []ImsiDigit
}
