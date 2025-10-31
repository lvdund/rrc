package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1700-broadcastSCell-r17 ::= ENUMERATED
type FeaturesetdownlinkperccV1700BroadcastscellR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccV1700BroadcastscellR17EnumeratedNothing = iota
	FeaturesetdownlinkperccV1700BroadcastscellR17EnumeratedSupported
)
