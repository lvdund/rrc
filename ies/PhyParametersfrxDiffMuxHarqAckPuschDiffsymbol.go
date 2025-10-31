package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-mux-HARQ-ACK-PUSCH-DiffSymbol ::= ENUMERATED
type PhyParametersfrxDiffMuxHarqAckPuschDiffsymbol struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMuxHarqAckPuschDiffsymbolEnumeratedNothing = iota
	PhyParametersfrxDiffMuxHarqAckPuschDiffsymbolEnumeratedSupported
)
