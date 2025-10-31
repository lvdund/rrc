package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot-diffSymbol ::= ENUMERATED
type PhyParametersfrxDiffMuxSrHarqAckCsiPucchOnceperslotDiffsymbol struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMuxSrHarqAckCsiPucchOnceperslotDiffsymbolEnumeratedNothing = iota
	PhyParametersfrxDiffMuxSrHarqAckCsiPucchOnceperslotDiffsymbolEnumeratedSupported
)
