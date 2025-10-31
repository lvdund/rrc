package ies

import "rrc/utils"

// Phy-ParametersCommon-twoHARQ-ACK-CodebookForUnicastAndMulticast-r17 ::= ENUMERATED
type PhyParameterscommonTwoharqAckCodebookforunicastandmulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonTwoharqAckCodebookforunicastandmulticastR17EnumeratedNothing = iota
	PhyParameterscommonTwoharqAckCodebookforunicastandmulticastR17EnumeratedSupported
)
