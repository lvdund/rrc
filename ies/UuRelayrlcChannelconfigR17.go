package ies

import "rrc/utils"

// Uu-RelayRLC-ChannelConfig-r17 ::= SEQUENCE
// Extensible
type UuRelayrlcChannelconfigR17 struct {
	UuLogicalchannelidentityR17 *Logicalchannelidentity
	UuRelayrlcChannelidR17      UuRelayrlcChannelidR17
	ReestablishrlcR17           *utils.BOOLEAN
	RlcConfigR17                *RlcConfig
	MacLogicalchannelconfigR17  *Logicalchannelconfig
}
