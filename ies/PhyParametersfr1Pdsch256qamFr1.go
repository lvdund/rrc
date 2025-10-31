package ies

import "rrc/utils"

// Phy-ParametersFR1-pdsch-256QAM-FR1 ::= ENUMERATED
type PhyParametersfr1Pdsch256qamFr1 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr1Pdsch256qamFr1EnumeratedNothing = iota
	PhyParametersfr1Pdsch256qamFr1EnumeratedSupported
)
