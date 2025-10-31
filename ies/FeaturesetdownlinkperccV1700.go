package ies

// FeatureSetDownlinkPerCC-v1700 ::= SEQUENCE
type FeaturesetdownlinkperccV1700 struct {
	SupportedminbandwidthdlR17           *SupportedbandwidthV1700
	BroadcastscellR17                    *FeaturesetdownlinkperccV1700BroadcastscellR17
	MaxnumbermimoLayersmulticastpdschR17 *FeaturesetdownlinkperccV1700MaxnumbermimoLayersmulticastpdschR17
	DynamicmulticastscellR17             *FeaturesetdownlinkperccV1700DynamicmulticastscellR17
	SupportedbandwidthdlV1710            *SupportedbandwidthV1700
	SupportedcrsInterfmitigationR17      *CrsInterfmitigationR17
}
