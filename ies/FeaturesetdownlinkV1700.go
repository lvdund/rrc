package ies

// FeatureSetDownlink-v1700 ::= SEQUENCE
type FeaturesetdownlinkV1700 struct {
	Scalingfactor1024qamFr1R17    *FeaturesetdownlinkV1700Scalingfactor1024qamFr1R17
	TimedurationforqclV1710       *FeaturesetdownlinkV1700TimedurationforqclV1710
	SfnSchemeaR17                 *FeaturesetdownlinkV1700SfnSchemeaR17
	SfnSchemeaPdcchOnlyR17        *FeaturesetdownlinkV1700SfnSchemeaPdcchOnlyR17
	SfnSchemeaDynamicswitchingR17 *FeaturesetdownlinkV1700SfnSchemeaDynamicswitchingR17
	SfnSchemeaPdschOnlyR17        *FeaturesetdownlinkV1700SfnSchemeaPdschOnlyR17
	SfnSchemebR17                 *FeaturesetdownlinkV1700SfnSchemebR17
	SfnSchemebDynamicswitchingR17 *FeaturesetdownlinkV1700SfnSchemebDynamicswitchingR17
	SfnSchemebPdschOnlyR17        *FeaturesetdownlinkV1700SfnSchemebPdschOnlyR17
	MtrpPdcchCase21spangapR17     *FeaturesetdownlinkV1700MtrpPdcchCase21spangapR17
	MtrpPdcchLegacymonitoringR17  *FeaturesetdownlinkV1700MtrpPdcchLegacymonitoringR17
	MtrpPdcchMultidciMultitrpR17  *FeaturesetdownlinkV1700MtrpPdcchMultidciMultitrpR17
	DynamicmulticastpcellR17      *FeaturesetdownlinkV1700DynamicmulticastpcellR17
	MtrpPdcchRepetitionR17        *FeaturesetdownlinkV1700MtrpPdcchRepetitionR17
}
