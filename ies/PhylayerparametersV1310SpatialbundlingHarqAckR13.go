package ies

import "rrc/utils"

// PhyLayerParameters-v1310-spatialBundling-HARQ-ACK-r13 ::= ENUMERATED
type PhylayerparametersV1310SpatialbundlingHarqAckR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310SpatialbundlingHarqAckR13EnumeratedNothing = iota
	PhylayerparametersV1310SpatialbundlingHarqAckR13EnumeratedSupported
)
