package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-dl-64QAM-MCS-TableAlt ::= ENUMERATED
type PhyParametersfrxDiffDl64qamMcsTablealt struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffDl64qamMcsTablealtEnumeratedNothing = iota
	PhyParametersfrxDiffDl64qamMcsTablealtEnumeratedSupported
)
