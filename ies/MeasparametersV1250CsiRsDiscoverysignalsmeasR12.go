package ies

import "rrc/utils"

// MeasParameters-v1250-csi-RS-DiscoverySignalsMeas-r12 ::= ENUMERATED
type MeasparametersV1250CsiRsDiscoverysignalsmeasR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250CsiRsDiscoverysignalsmeasR12EnumeratedNothing = iota
	MeasparametersV1250CsiRsDiscoverysignalsmeasR12EnumeratedSupported
)
