package ies

// VarMeasConfig-speedStatePars ::= CHOICE
const (
	VarmeasconfigSpeedstateparsChoiceNothing = iota
	VarmeasconfigSpeedstateparsChoiceRelease
	VarmeasconfigSpeedstateparsChoiceSetup
)

type VarmeasconfigSpeedstatepars struct {
	Choice  uint64
	Release *struct{}
	Setup   *VarmeasconfigSpeedstateparsSetup
}
