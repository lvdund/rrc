package ies

import "rrc/utils"

// LogicalChannelConfig-bitRateMultiplier-r16 ::= ENUMERATED
type LogicalchannelconfigBitratemultiplierR16 struct {
	Value utils.ENUMERATED
}

const (
	LogicalchannelconfigBitratemultiplierR16EnumeratedNothing = iota
	LogicalchannelconfigBitratemultiplierR16EnumeratedX40
	LogicalchannelconfigBitratemultiplierR16EnumeratedX70
	LogicalchannelconfigBitratemultiplierR16EnumeratedX100
	LogicalchannelconfigBitratemultiplierR16EnumeratedX200
)
