package ies

import "rrc/utils"

// MobilityParameters-v1610 ::= SEQUENCE
type MobilityparametersV1610 struct {
	ChoR16                 *utils.ENUMERATED
	ChoFddTddR16           *utils.ENUMERATED
	ChoFailureR16          *utils.ENUMERATED
	ChoTwotriggereventsR16 *utils.ENUMERATED
}
