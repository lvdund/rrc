package ies

import "rrc/utils"

// FeatureSetDownlink-v1700-dynamicMulticastPCell-r17 ::= ENUMERATED
type FeaturesetdownlinkV1700DynamicmulticastpcellR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1700DynamicmulticastpcellR17EnumeratedNothing = iota
	FeaturesetdownlinkV1700DynamicmulticastpcellR17EnumeratedSupported
)
