package ies

import "rrc/utils"

// PhyLayerParameters-v1250-discoverySignalsInDeactSCell-r12 ::= ENUMERATED
type PhylayerparametersV1250DiscoverysignalsindeactscellR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250DiscoverysignalsindeactscellR12EnumeratedNothing = iota
	PhylayerparametersV1250DiscoverysignalsindeactscellR12EnumeratedSupported
)
