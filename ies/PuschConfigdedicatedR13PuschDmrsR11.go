package ies

// PUSCH-ConfigDedicated-r13-pusch-DMRS-r11 ::= CHOICE
const (
	PuschConfigdedicatedR13PuschDmrsR11ChoiceNothing = iota
	PuschConfigdedicatedR13PuschDmrsR11ChoiceRelease
	PuschConfigdedicatedR13PuschDmrsR11ChoiceSetup
)

type PuschConfigdedicatedR13PuschDmrsR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschConfigdedicatedR13PuschDmrsR11Setup
}
