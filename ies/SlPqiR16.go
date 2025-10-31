package ies

import "rrc/utils"

// SL-PQI-r16 ::= CHOICE
// Extensible
const (
	SlPqiR16ChoiceNothing = iota
	SlPqiR16ChoiceSlStandardizedpqiR16
	SlPqiR16ChoiceSlNonStandardizedpqiR16
)

type SlPqiR16 struct {
	Choice                  uint64
	SlStandardizedpqiR16    *utils.INTEGER `lb:0,ub:255`
	SlNonStandardizedpqiR16 *SlPqiR16SlNonStandardizedpqiR16
}
