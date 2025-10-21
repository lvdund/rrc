package ies

import "rrc/utils"

// FeatureSetsEUTRA-r15 ::= SEQUENCE
// Extensible
type FeaturesetseutraR15 struct {
	FeaturesetsdlR15      *interface{}
	FeaturesetsdlPerccR15 *interface{}
	FeaturesetsulR15      *interface{}
	FeaturesetsulPerccR15 *interface{}
}
