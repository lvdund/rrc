package ies

import "rrc/utils"

// FeatureSetDownlink-v1720-rtt-BasedPDC-CSI-RS-ForTracking-r17 ::= ENUMERATED
type FeaturesetdownlinkV1720RttBasedpdcCsiRsFortrackingR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1720RttBasedpdcCsiRsFortrackingR17EnumeratedNothing = iota
	FeaturesetdownlinkV1720RttBasedpdcCsiRsFortrackingR17EnumeratedSupported
)
