package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-mux-SR-HARQ-ACK-PUCCH ::= ENUMERATED
type PhyParametersfrxDiffMuxSrHarqAckPucch struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMuxSrHarqAckPucchEnumeratedNothing = iota
	PhyParametersfrxDiffMuxSrHarqAckPucchEnumeratedSupported
)
