package ies

import "rrc/utils"

// Phy-ParametersCommon-dynamicHARQ-ACK-CodeB-CBG-Retx-DL ::= ENUMERATED
type PhyParameterscommonDynamicharqAckCodebCbgRetxDl struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDynamicharqAckCodebCbgRetxDlEnumeratedNothing = iota
	PhyParameterscommonDynamicharqAckCodebCbgRetxDlEnumeratedSupported
)
