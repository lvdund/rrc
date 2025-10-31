package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot ::= ENUMERATED
type PhyParametersfrxDiffMuxSrHarqAckCsiPucchMultiperslot struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMuxSrHarqAckCsiPucchMultiperslotEnumeratedNothing = iota
	PhyParametersfrxDiffMuxSrHarqAckCsiPucchMultiperslotEnumeratedSupported
)
