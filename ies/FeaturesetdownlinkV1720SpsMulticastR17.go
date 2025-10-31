package ies

import "rrc/utils"

// FeatureSetDownlink-v1720-sps-Multicast-r17 ::= ENUMERATED
type FeaturesetdownlinkV1720SpsMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1720SpsMulticastR17EnumeratedNothing = iota
	FeaturesetdownlinkV1720SpsMulticastR17EnumeratedSupported
)
