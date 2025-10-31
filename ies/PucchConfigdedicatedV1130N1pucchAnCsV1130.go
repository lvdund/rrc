package ies

// PUCCH-ConfigDedicated-v1130-n1PUCCH-AN-CS-v1130 ::= CHOICE
const (
	PucchConfigdedicatedV1130N1pucchAnCsV1130ChoiceNothing = iota
	PucchConfigdedicatedV1130N1pucchAnCsV1130ChoiceRelease
	PucchConfigdedicatedV1130N1pucchAnCsV1130ChoiceSetup
)

type PucchConfigdedicatedV1130N1pucchAnCsV1130 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchConfigdedicatedV1130N1pucchAnCsV1130Setup
}
