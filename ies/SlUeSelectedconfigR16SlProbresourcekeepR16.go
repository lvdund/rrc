package ies

import "rrc/utils"

// SL-UE-SelectedConfig-r16-sl-ProbResourceKeep-r16 ::= ENUMERATED
type SlUeSelectedconfigR16SlProbresourcekeepR16 struct {
	Value utils.ENUMERATED
}

const (
	SlUeSelectedconfigR16SlProbresourcekeepR16EnumeratedNothing = iota
	SlUeSelectedconfigR16SlProbresourcekeepR16EnumeratedV0
	SlUeSelectedconfigR16SlProbresourcekeepR16EnumeratedV0dot2
	SlUeSelectedconfigR16SlProbresourcekeepR16EnumeratedV0dot4
	SlUeSelectedconfigR16SlProbresourcekeepR16EnumeratedV0dot6
	SlUeSelectedconfigR16SlProbresourcekeepR16EnumeratedV0dot8
)
