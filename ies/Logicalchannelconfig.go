package ies

import "rrc/utils"

// LogicalChannelConfig ::= SEQUENCE
// Extensible
type Logicalchannelconfig struct {
	UlSpecificparameters     *LogicalchannelconfigUlSpecificparameters `ext`
	ChannelaccesspriorityR16 *utils.INTEGER                            `lb:0,ub:4,ext`
	BitratemultiplierR16     *LogicalchannelconfigBitratemultiplierR16 `ext`
}
