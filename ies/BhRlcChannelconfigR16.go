package ies

import "rrc/utils"

// BH-RLC-ChannelConfig-r16 ::= SEQUENCE
// Extensible
type BhRlcChannelconfigR16 struct {
	BhLogicalchannelidentityR16 *BhLogicalchannelidentityR16
	BhRlcChannelidR16           BhRlcChannelidR16
	ReestablishrlcR16           *utils.BOOLEAN
	RlcConfigR16                *RlcConfig
	MacLogicalchannelconfigR16  *Logicalchannelconfig
}
