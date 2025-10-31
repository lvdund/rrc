package ies

import "rrc/utils"

// MAC-ParametersCommon-lch-PriorityBasedPrioritization-r16 ::= ENUMERATED
type MacParameterscommonLchPrioritybasedprioritizationR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLchPrioritybasedprioritizationR16EnumeratedNothing = iota
	MacParameterscommonLchPrioritybasedprioritizationR16EnumeratedSupported
)
