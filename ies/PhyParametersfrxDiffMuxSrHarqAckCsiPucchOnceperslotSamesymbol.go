package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot-sameSymbol ::= ENUMERATED
type PhyParametersfrxDiffMuxSrHarqAckCsiPucchOnceperslotSamesymbol struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMuxSrHarqAckCsiPucchOnceperslotSamesymbolEnumeratedNothing = iota
	PhyParametersfrxDiffMuxSrHarqAckCsiPucchOnceperslotSamesymbolEnumeratedSupported
)
