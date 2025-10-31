package ies

// MeasGapConfigPerCC-List-r14 ::= CHOICE
const (
	MeasgapconfigperccListR14ChoiceNothing = iota
	MeasgapconfigperccListR14ChoiceRelease
	MeasgapconfigperccListR14ChoiceSetup
)

type MeasgapconfigperccListR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeasgapconfigperccListR14Setup
}
