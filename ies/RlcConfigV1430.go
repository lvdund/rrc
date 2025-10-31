package ies

// RLC-Config-v1430 ::= CHOICE
const (
	RlcConfigV1430ChoiceNothing = iota
	RlcConfigV1430ChoiceRelease
	RlcConfigV1430ChoiceSetup
)

type RlcConfigV1430 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlcConfigV1430Setup
}
