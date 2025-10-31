package ies

// FeatureSetsEUTRA-r15 ::= SEQUENCE
// Extensible
type FeaturesetseutraR15 struct {
	FeaturesetsdlR15      *[]FeaturesetdlR15      `lb:1,ub:maxFeatureSetsR15`
	FeaturesetsdlPerccR15 *[]FeaturesetdlPerccR15 `lb:1,ub:maxPerCCFeaturesetsR15`
	FeaturesetsulR15      *[]FeaturesetulR15      `lb:1,ub:maxFeatureSetsR15`
	FeaturesetsulPerccR15 *[]FeaturesetulPerccR15 `lb:1,ub:maxPerCCFeaturesetsR15`
}
