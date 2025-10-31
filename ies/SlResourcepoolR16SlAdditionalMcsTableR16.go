package ies

import "rrc/utils"

// SL-ResourcePool-r16-sl-Additional-MCS-Table-r16 ::= ENUMERATED
type SlResourcepoolR16SlAdditionalMcsTableR16 struct {
	Value utils.ENUMERATED
}

const (
	SlResourcepoolR16SlAdditionalMcsTableR16EnumeratedNothing = iota
	SlResourcepoolR16SlAdditionalMcsTableR16EnumeratedQam256
	SlResourcepoolR16SlAdditionalMcsTableR16EnumeratedQam64lowse
	SlResourcepoolR16SlAdditionalMcsTableR16EnumeratedQam256_Qam64lowse
)
