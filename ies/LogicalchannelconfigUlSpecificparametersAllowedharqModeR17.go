package ies

import "rrc/utils"

// LogicalChannelConfig-ul-SpecificParameters-allowedHARQ-mode-r17 ::= ENUMERATED
type LogicalchannelconfigUlSpecificparametersAllowedharqModeR17 struct {
	Value utils.ENUMERATED
}

const (
	LogicalchannelconfigUlSpecificparametersAllowedharqModeR17EnumeratedNothing = iota
	LogicalchannelconfigUlSpecificparametersAllowedharqModeR17EnumeratedHarqmodea
	LogicalchannelconfigUlSpecificparametersAllowedharqModeR17EnumeratedHarqmodeb
)
