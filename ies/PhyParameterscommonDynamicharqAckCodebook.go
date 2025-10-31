package ies

import "rrc/utils"

// Phy-ParametersCommon-dynamicHARQ-ACK-Codebook ::= ENUMERATED
type PhyParameterscommonDynamicharqAckCodebook struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDynamicharqAckCodebookEnumeratedNothing = iota
	PhyParameterscommonDynamicharqAckCodebookEnumeratedSupported
)
