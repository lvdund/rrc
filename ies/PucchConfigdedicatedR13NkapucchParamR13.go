package ies

// PUCCH-ConfigDedicated-r13-nkaPUCCH-Param-r13 ::= CHOICE
const (
	PucchConfigdedicatedR13NkapucchParamR13ChoiceNothing = iota
	PucchConfigdedicatedR13NkapucchParamR13ChoiceRelease
	PucchConfigdedicatedR13NkapucchParamR13ChoiceSetup
)

type PucchConfigdedicatedR13NkapucchParamR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedR13NkapucchParamR13Setup
}
