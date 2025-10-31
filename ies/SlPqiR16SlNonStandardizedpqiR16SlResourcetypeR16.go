package ies

import "rrc/utils"

// SL-PQI-r16-sl-Non-StandardizedPQI-r16-sl-ResourceType-r16 ::= ENUMERATED
type SlPqiR16SlNonStandardizedpqiR16SlResourcetypeR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPqiR16SlNonStandardizedpqiR16SlResourcetypeR16EnumeratedNothing = iota
	SlPqiR16SlNonStandardizedpqiR16SlResourcetypeR16EnumeratedGbr
	SlPqiR16SlNonStandardizedpqiR16SlResourcetypeR16EnumeratedNon_Gbr
	SlPqiR16SlNonStandardizedpqiR16SlResourcetypeR16EnumeratedDelaycriticalgbr
	SlPqiR16SlNonStandardizedpqiR16SlResourcetypeR16EnumeratedSpare1
)
