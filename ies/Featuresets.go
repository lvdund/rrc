package ies

// FeatureSets ::= SEQUENCE
// Extensible
type Featuresets struct {
	Featuresetsdownlink           *[]Featuresetdownlink           `lb:1,ub:maxDownlinkFeatureSets`
	Featuresetsdownlinkpercc      *[]Featuresetdownlinkpercc      `lb:1,ub:maxPerCCFeaturesets`
	Featuresetsuplink             *[]Featuresetuplink             `lb:1,ub:maxUplinkFeatureSets`
	Featuresetsuplinkpercc        *[]Featuresetuplinkpercc        `lb:1,ub:maxPerCCFeaturesets`
	FeaturesetsdownlinkV1540      *[]FeaturesetdownlinkV1540      `lb:1,ub:maxDownlinkFeatureSets,ext`
	FeaturesetsuplinkV1540        *[]FeaturesetuplinkV1540        `lb:1,ub:maxUplinkFeatureSets,ext`
	FeaturesetsuplinkperccV1540   *[]FeaturesetuplinkperccV1540   `lb:1,ub:maxPerCCFeaturesets,ext`
	FeaturesetsdownlinkV15a0      *[]FeaturesetdownlinkV15a0      `lb:1,ub:maxDownlinkFeatureSets,ext`
	FeaturesetsdownlinkV1610      *[]FeaturesetdownlinkV1610      `lb:1,ub:maxDownlinkFeatureSets,ext`
	FeaturesetsuplinkV1610        *[]FeaturesetuplinkV1610        `lb:1,ub:maxUplinkFeatureSets,ext`
	FeaturesetdownlinkperccV1620  *[]FeaturesetdownlinkperccV1620 `lb:1,ub:maxPerCCFeaturesets,ext`
	FeaturesetsuplinkV1630        *[]FeaturesetuplinkV1630        `lb:1,ub:maxUplinkFeatureSets,ext`
	FeaturesetsuplinkV1640        *[]FeaturesetuplinkV1640        `lb:1,ub:maxUplinkFeatureSets,ext`
	FeaturesetsdownlinkV1700      *[]FeaturesetdownlinkV1700      `lb:1,ub:maxDownlinkFeatureSets,ext`
	FeaturesetsdownlinkperccV1700 *[]FeaturesetdownlinkperccV1700 `lb:1,ub:maxPerCCFeaturesets,ext`
	FeaturesetsuplinkV1710        *[]FeaturesetuplinkV1710        `lb:1,ub:maxUplinkFeatureSets,ext`
	FeaturesetsuplinkperccV1700   *[]FeaturesetuplinkperccV1700   `lb:1,ub:maxPerCCFeaturesets,ext`
	FeaturesetsdownlinkV1720      *[]FeaturesetdownlinkV1720      `lb:1,ub:maxDownlinkFeatureSets,ext`
	FeaturesetsdownlinkperccV1720 *[]FeaturesetdownlinkperccV1720 `lb:1,ub:maxPerCCFeaturesets,ext`
	FeaturesetsuplinkV1720        *[]FeaturesetuplinkV1720        `lb:1,ub:maxUplinkFeatureSets,ext`
	FeaturesetsdownlinkV1730      *[]FeaturesetdownlinkV1730      `lb:1,ub:maxDownlinkFeatureSets,ext`
	FeaturesetsdownlinkperccV1730 *[]FeaturesetdownlinkperccV1730 `lb:1,ub:maxPerCCFeaturesets,ext`
}
