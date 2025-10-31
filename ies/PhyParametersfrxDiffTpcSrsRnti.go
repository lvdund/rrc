package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-tpc-SRS-RNTI ::= ENUMERATED
type PhyParametersfrxDiffTpcSrsRnti struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTpcSrsRntiEnumeratedNothing = iota
	PhyParametersfrxDiffTpcSrsRntiEnumeratedSupported
)
