package ies

import "rrc/utils"

// N1PUCCH-AN-CS-r10 ::= SEQUENCE OF INTEGER
// SIZE (1..4)
type N1pucchAnCsR10 struct {
	Value []Integer
}
