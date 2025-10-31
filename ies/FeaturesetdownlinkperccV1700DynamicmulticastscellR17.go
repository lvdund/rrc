package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1700-dynamicMulticastSCell-r17 ::= ENUMERATED
type FeaturesetdownlinkperccV1700DynamicmulticastscellR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccV1700DynamicmulticastscellR17EnumeratedNothing = iota
	FeaturesetdownlinkperccV1700DynamicmulticastscellR17EnumeratedSupported
)
