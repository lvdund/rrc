package ies

import "rrc/utils"

// MeasParameters-v1250-crs-DiscoverySignalsMeas-r12 ::= ENUMERATED
type MeasparametersV1250CrsDiscoverysignalsmeasR12 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1250CrsDiscoverysignalsmeasR12EnumeratedNothing = iota
	MeasparametersV1250CrsDiscoverysignalsmeasR12EnumeratedSupported
)
