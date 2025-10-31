package ies

import "rrc/utils"

// Phy-ParametersCommon-spatialBundlingHARQ-ACK ::= ENUMERATED
type PhyParameterscommonSpatialbundlingharqAck struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSpatialbundlingharqAckEnumeratedNothing = iota
	PhyParameterscommonSpatialbundlingharqAckEnumeratedSupported
)
