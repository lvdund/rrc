package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-cqi-TableAlt ::= ENUMERATED
type PhyParametersfrxDiffCqiTablealt struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffCqiTablealtEnumeratedNothing = iota
	PhyParametersfrxDiffCqiTablealtEnumeratedSupported
)
