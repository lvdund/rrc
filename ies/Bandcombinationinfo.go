package ies

// BandCombinationInfo ::= SEQUENCE
type Bandcombinationinfo struct {
	Bandcombinationindex   Bandcombinationindex
	Allowedfeaturesetslist []Featuresetentryindex `lb:1,ub:maxFeatureSetsPerBand`
}
