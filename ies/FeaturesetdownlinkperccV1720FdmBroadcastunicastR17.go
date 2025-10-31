package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1720-fdm-BroadcastUnicast-r17 ::= ENUMERATED
type FeaturesetdownlinkperccV1720FdmBroadcastunicastR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccV1720FdmBroadcastunicastR17EnumeratedNothing = iota
	FeaturesetdownlinkperccV1720FdmBroadcastunicastR17EnumeratedSupported
)
