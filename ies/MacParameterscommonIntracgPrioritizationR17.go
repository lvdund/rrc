package ies

import "rrc/utils"

// MAC-ParametersCommon-intraCG-Prioritization-r17 ::= ENUMERATED
type MacParameterscommonIntracgPrioritizationR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonIntracgPrioritizationR17EnumeratedNothing = iota
	MacParameterscommonIntracgPrioritizationR17EnumeratedSupported
)
