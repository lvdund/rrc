package ies

// MeasSubframePatternConfigNeigh-r10 ::= CHOICE
const (
	MeassubframepatternconfigneighR10ChoiceNothing = iota
	MeassubframepatternconfigneighR10ChoiceRelease
	MeassubframepatternconfigneighR10ChoiceSetup
)

type MeassubframepatternconfigneighR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeassubframepatternconfigneighR10Setup
}
