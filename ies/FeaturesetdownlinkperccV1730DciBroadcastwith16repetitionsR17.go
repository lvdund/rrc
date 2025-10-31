package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-v1730-dci-BroadcastWith16Repetitions-r17 ::= ENUMERATED
type FeaturesetdownlinkperccV1730DciBroadcastwith16repetitionsR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccV1730DciBroadcastwith16repetitionsR17EnumeratedNothing = iota
	FeaturesetdownlinkperccV1730DciBroadcastwith16repetitionsR17EnumeratedSupported
)
