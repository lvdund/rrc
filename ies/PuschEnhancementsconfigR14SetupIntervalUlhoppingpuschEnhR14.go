package ies

import "rrc/utils"

// PUSCH-EnhancementsConfig-r14-setup-interval-ULHoppingPUSCH-Enh-r14 ::= CHOICE
const (
	PuschEnhancementsconfigR14SetupIntervalUlhoppingpuschEnhR14ChoiceNothing = iota
	PuschEnhancementsconfigR14SetupIntervalUlhoppingpuschEnhR14ChoiceIntervalFddPuschEnhR14
	PuschEnhancementsconfigR14SetupIntervalUlhoppingpuschEnhR14ChoiceIntervalTddPuschEnhR14
)

type PuschEnhancementsconfigR14SetupIntervalUlhoppingpuschEnhR14 struct {
	Choice                 uint64
	IntervalFddPuschEnhR14 *utils.ENUMERATED
	IntervalTddPuschEnhR14 *utils.ENUMERATED
}
