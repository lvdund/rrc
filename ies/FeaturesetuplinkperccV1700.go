package ies

// FeatureSetUplinkPerCC-v1700 ::= SEQUENCE
type FeaturesetuplinkperccV1700 struct {
	SupportedminbandwidthulR17  *SupportedbandwidthV1700
	MtrpPuschRepetitiontypebR17 *FeaturesetuplinkperccV1700MtrpPuschRepetitiontypebR17
	MtrpPuschTypebCbR17         *FeaturesetuplinkperccV1700MtrpPuschTypebCbR17
	SupportedbandwidthulV1710   *SupportedbandwidthV1700
}
