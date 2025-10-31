package ies

// PUCCH-ConfigDedicated-v1130-nPUCCH-Param-r11 ::= CHOICE
const (
	PucchConfigdedicatedV1130NpucchParamR11ChoiceNothing = iota
	PucchConfigdedicatedV1130NpucchParamR11ChoiceRelease
	PucchConfigdedicatedV1130NpucchParamR11ChoiceSetup
)

type PucchConfigdedicatedV1130NpucchParamR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedV1130NpucchParamR11Setup
}
