package ies

import "rrc/utils"

// Phy-ParametersCommon-semiStaticHARQ-ACK-Codebook ::= ENUMERATED
type PhyParameterscommonSemistaticharqAckCodebook struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSemistaticharqAckCodebookEnumeratedNothing = iota
	PhyParameterscommonSemistaticharqAckCodebookEnumeratedSupported
)
