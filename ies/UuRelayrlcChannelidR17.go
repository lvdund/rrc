package ies

import "rrc/utils"

// Uu-RelayRLC-ChannelID-r17 ::= utils.INTEGER (1..maxLC-ID)
type UuRelayrlcChannelidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxLCId`
}
