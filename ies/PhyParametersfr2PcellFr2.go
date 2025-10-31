package ies

import "rrc/utils"

// Phy-ParametersFR2-pCell-FR2 ::= ENUMERATED
type PhyParametersfr2PcellFr2 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr2PcellFr2EnumeratedNothing = iota
	PhyParametersfr2PcellFr2EnumeratedSupported
)
