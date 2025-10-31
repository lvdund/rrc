package ies

import "rrc/utils"

// Phy-ParametersCommon-multiPUCCH-HARQ-ACK-ForMulticastUnicast-r17 ::= ENUMERATED
type PhyParameterscommonMultipucchHarqAckFormulticastunicastR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonMultipucchHarqAckFormulticastunicastR17EnumeratedNothing = iota
	PhyParameterscommonMultipucchHarqAckFormulticastunicastR17EnumeratedSupported
)
