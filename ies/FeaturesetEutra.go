package ies

// FeatureSet-eutra ::= SEQUENCE
type FeaturesetEutra struct {
	Downlinkseteutra FeatureseteutraDownlinkid
	Uplinkseteutra   FeatureseteutraUplinkid
}
