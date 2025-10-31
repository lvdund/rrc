package ies

// PUCCH-ConfigDedicated-v1370-pucch-Format-v1370 ::= CHOICE
const (
	PucchConfigdedicatedV1370PucchFormatV1370ChoiceNothing = iota
	PucchConfigdedicatedV1370PucchFormatV1370ChoiceRelease
	PucchConfigdedicatedV1370PucchFormatV1370ChoiceSetup
)

type PucchConfigdedicatedV1370PucchFormatV1370 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchFormat3ConfR13
}
