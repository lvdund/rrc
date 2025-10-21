package ies

import "rrc/utils"

// MobilityParameters-r14 ::= SEQUENCE
type MobilityparametersR14 struct {
	MakebeforebreakR14 *utils.ENUMERATED
	RachLessR14        *utils.ENUMERATED
}
