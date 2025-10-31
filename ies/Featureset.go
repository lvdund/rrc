package ies

// FeatureSet ::= CHOICE
const (
	FeaturesetChoiceNothing = iota
	FeaturesetChoiceEutra
	FeaturesetChoiceNr
)

type Featureset struct {
	Choice uint64
	Eutra  *FeaturesetEutra
	Nr     *FeaturesetNr
}
