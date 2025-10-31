package ies

// PUCCH-ConfigDedicated-v1250-nkaPUCCH-Param-r12 ::= CHOICE
const (
	PucchConfigdedicatedV1250NkapucchParamR12ChoiceNothing = iota
	PucchConfigdedicatedV1250NkapucchParamR12ChoiceRelease
	PucchConfigdedicatedV1250NkapucchParamR12ChoiceSetup
)

type PucchConfigdedicatedV1250NkapucchParamR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedV1250NkapucchParamR12Setup
}
