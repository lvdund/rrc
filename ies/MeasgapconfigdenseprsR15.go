package ies

// MeasGapConfigDensePRS-r15 ::= CHOICE
// Extensible
const (
	MeasgapconfigdenseprsR15ChoiceNothing = iota
	MeasgapconfigdenseprsR15ChoiceRelease
	MeasgapconfigdenseprsR15ChoiceSetup
)

type MeasgapconfigdenseprsR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeasgapconfigdenseprsR15Setup
}
