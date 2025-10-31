package ies

import "rrc/utils"

// BandParametersSidelinkDiscovery-r17-sl-CrossCarrierScheduling-r17 ::= ENUMERATED
type BandparameterssidelinkdiscoveryR17SlCrosscarrierschedulingR17 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkdiscoveryR17SlCrosscarrierschedulingR17EnumeratedNothing = iota
	BandparameterssidelinkdiscoveryR17SlCrosscarrierschedulingR17EnumeratedSupported
)
