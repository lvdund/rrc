package ies

// SPUCCH-Config-v1550 ::= CHOICE
const (
	SpucchConfigV1550ChoiceNothing = iota
	SpucchConfigV1550ChoiceRelease
	SpucchConfigV1550ChoiceSetup
)

type SpucchConfigV1550 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpucchConfigV1550Setup
}
