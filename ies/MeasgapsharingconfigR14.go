package ies

// MeasGapSharingConfig-r14 ::= CHOICE
const (
	MeasgapsharingconfigR14ChoiceNothing = iota
	MeasgapsharingconfigR14ChoiceRelease
	MeasgapsharingconfigR14ChoiceSetup
)

type MeasgapsharingconfigR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeasgapsharingconfigR14Setup
}
