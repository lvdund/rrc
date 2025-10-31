package ies

import "rrc/utils"

// HighSpeedParameters-v1650 ::= CHOICE
const (
	HighspeedparametersV1650ChoiceNothing = iota
	HighspeedparametersV1650ChoiceIntranrMeasurementenhancementR16
	HighspeedparametersV1650ChoiceInterratMeasurementenhancementR16
)

type HighspeedparametersV1650 struct {
	Choice                            uint64
	IntranrMeasurementenhancementR16  *utils.ENUMERATED
	InterratMeasurementenhancementR16 *utils.ENUMERATED
}
