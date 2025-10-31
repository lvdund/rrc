package ies

import "rrc/utils"

// Phy-ParametersMRDC-fdd-PCellUL-TX-AllUL-Subframe-r16 ::= ENUMERATED
type PhyParametersmrdcFddPcellulTxAllulSubframeR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersmrdcFddPcellulTxAllulSubframeR16EnumeratedNothing = iota
	PhyParametersmrdcFddPcellulTxAllulSubframeR16EnumeratedSupported
)
