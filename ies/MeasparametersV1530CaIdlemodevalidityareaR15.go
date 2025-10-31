package ies

import "rrc/utils"

// MeasParameters-v1530-ca-IdleModeValidityArea-r15 ::= ENUMERATED
type MeasparametersV1530CaIdlemodevalidityareaR15 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1530CaIdlemodevalidityareaR15EnumeratedNothing = iota
	MeasparametersV1530CaIdlemodevalidityareaR15EnumeratedSupported
)
