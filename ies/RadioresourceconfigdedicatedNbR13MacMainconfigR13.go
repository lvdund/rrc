package ies

// RadioResourceConfigDedicated-NB-r13-mac-MainConfig-r13 ::= CHOICE
const (
	RadioresourceconfigdedicatedNbR13MacMainconfigR13ChoiceNothing = iota
	RadioresourceconfigdedicatedNbR13MacMainconfigR13ChoiceExplicitvalueR13
	RadioresourceconfigdedicatedNbR13MacMainconfigR13ChoiceDefaultvalueR13
)

type RadioresourceconfigdedicatedNbR13MacMainconfigR13 struct {
	Choice           uint64
	ExplicitvalueR13 *MacMainconfigNbR13
	DefaultvalueR13  *struct{}
}
