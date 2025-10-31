package ies

// FeatureSetCombination ::= SEQUENCE OF FeatureSetsPerBand
// SIZE (1..maxSimultaneousBands)
type Featuresetcombination struct {
	Value []Featuresetsperband `lb:1,ub:maxSimultaneousBands`
}
