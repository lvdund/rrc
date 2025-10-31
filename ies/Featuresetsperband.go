package ies

// FeatureSetsPerBand ::= SEQUENCE OF FeatureSet
// SIZE (1..maxFeatureSetsPerBand)
type Featuresetsperband struct {
	Value []Featureset `lb:1,ub:maxFeatureSetsPerBand`
}
