package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1720-fdm-MulticastUnicast-r17 ::= ENUMERATED
type FeaturesetdownlinkperccV1720FdmMulticastunicastR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccV1720FdmMulticastunicastR17EnumeratedNothing = iota
	FeaturesetdownlinkperccV1720FdmMulticastunicastR17EnumeratedSupported
)
