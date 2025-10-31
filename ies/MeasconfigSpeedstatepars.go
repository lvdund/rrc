package ies

// MeasConfig-speedStatePars ::= CHOICE
const (
	MeasconfigSpeedstateparsChoiceNothing = iota
	MeasconfigSpeedstateparsChoiceRelease
	MeasconfigSpeedstateparsChoiceSetup
)

type MeasconfigSpeedstatepars struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeasconfigSpeedstateparsSetup
}
