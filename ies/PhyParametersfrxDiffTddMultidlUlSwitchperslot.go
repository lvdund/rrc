package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-tdd-MultiDL-UL-SwitchPerSlot ::= ENUMERATED
type PhyParametersfrxDiffTddMultidlUlSwitchperslot struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTddMultidlUlSwitchperslotEnumeratedNothing = iota
	PhyParametersfrxDiffTddMultidlUlSwitchperslotEnumeratedSupported
)
