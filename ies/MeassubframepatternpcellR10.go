package ies

// MeasSubframePatternPCell-r10 ::= CHOICE
const (
	MeassubframepatternpcellR10ChoiceNothing = iota
	MeassubframepatternpcellR10ChoiceRelease
	MeassubframepatternpcellR10ChoiceSetup
)

type MeassubframepatternpcellR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeassubframepatternR10
}
