package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-HARQ-AckBundling-r14 ::= ENUMERATED
type PhylayerparametersV1430CeHarqAckbundlingR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CeHarqAckbundlingR14EnumeratedNothing = iota
	PhylayerparametersV1430CeHarqAckbundlingR14EnumeratedSupported
)
