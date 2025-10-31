package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-tpc-PUSCH-RNTI ::= ENUMERATED
type PhyParametersfrxDiffTpcPuschRnti struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTpcPuschRntiEnumeratedNothing = iota
	PhyParametersfrxDiffTpcPuschRntiEnumeratedSupported
)
