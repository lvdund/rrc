package ies

// PUCCH-ConfigDedicated-r13-nPUCCH-Param-r13 ::= CHOICE
const (
	PucchConfigdedicatedR13NpucchParamR13ChoiceNothing = iota
	PucchConfigdedicatedR13NpucchParamR13ChoiceRelease
	PucchConfigdedicatedR13NpucchParamR13ChoiceSetup
)

type PucchConfigdedicatedR13NpucchParamR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedR13NpucchParamR13Setup
}
