package ies

import "rrc/utils"

// SL-RLC-ChannelConfig-r17 ::= SEQUENCE
// Extensible
type SlRlcChannelconfigR17 struct {
	SlRlcChannelidR17            SlRlcChannelidR17
	SlRlcConfigR17               *SlRlcConfigR16
	SlMacLogicalchannelconfigR17 *SlLogicalchannelconfigR16
	SlPacketdelaybudgetR17       *utils.INTEGER `lb:0,ub:1023`
}
