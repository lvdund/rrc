package ies

import "rrc/utils"

// Phy-ParametersMRDC-tdd-PCellUL-TX-AllUL-Subframe-r16 ::= ENUMERATED
type PhyParametersmrdcTddPcellulTxAllulSubframeR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersmrdcTddPcellulTxAllulSubframeR16EnumeratedNothing = iota
	PhyParametersmrdcTddPcellulTxAllulSubframeR16EnumeratedSupported
)
