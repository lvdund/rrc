package ies

// PUSCH-ConfigDedicated-v1130-pusch-DMRS-r11 ::= CHOICE
const (
	PuschConfigdedicatedV1130PuschDmrsR11ChoiceNothing = iota
	PuschConfigdedicatedV1130PuschDmrsR11ChoiceRelease
	PuschConfigdedicatedV1130PuschDmrsR11ChoiceSetup
)

type PuschConfigdedicatedV1130PuschDmrsR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschConfigdedicatedV1130PuschDmrsR11Setup
}
