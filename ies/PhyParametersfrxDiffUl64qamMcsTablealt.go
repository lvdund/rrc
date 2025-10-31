package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-ul-64QAM-MCS-TableAlt ::= ENUMERATED
type PhyParametersfrxDiffUl64qamMcsTablealt struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffUl64qamMcsTablealtEnumeratedNothing = iota
	PhyParametersfrxDiffUl64qamMcsTablealtEnumeratedSupported
)
