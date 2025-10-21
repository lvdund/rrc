package ies

import "rrc/utils"

// SL-PreconfigCommPoolList4-r12 ::= SEQUENCE OF SL-PreconfigCommPool-r12
// SIZE (1..maxSL-TxPool-r12)
type SlPreconfigcommpoollist4R12 struct {
	Value []SlPreconfigcommpoolR12
}
