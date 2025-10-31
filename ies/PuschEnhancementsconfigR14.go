package ies

// PUSCH-EnhancementsConfig-r14 ::= CHOICE
const (
	PuschEnhancementsconfigR14ChoiceNothing = iota
	PuschEnhancementsconfigR14ChoiceRelease
	PuschEnhancementsconfigR14ChoiceSetup
)

type PuschEnhancementsconfigR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschEnhancementsconfigR14Setup
}
