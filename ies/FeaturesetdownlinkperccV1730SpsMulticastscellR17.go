package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1730-sps-MulticastSCell-r17 ::= ENUMERATED
type FeaturesetdownlinkperccV1730SpsMulticastscellR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccV1730SpsMulticastscellR17EnumeratedNothing = iota
	FeaturesetdownlinkperccV1730SpsMulticastscellR17EnumeratedSupported
)
