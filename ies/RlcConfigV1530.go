package ies

// RLC-Config-v1530 ::= CHOICE
const (
	RlcConfigV1530ChoiceNothing = iota
	RlcConfigV1530ChoiceRelease
	RlcConfigV1530ChoiceSetup
)

type RlcConfigV1530 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlcConfigV1530Setup
}
