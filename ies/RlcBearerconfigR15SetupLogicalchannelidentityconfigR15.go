package ies

import "rrc/utils"

// RLC-BearerConfig-r15-setup-logicalChannelIdentityConfig-r15 ::= CHOICE
const (
	RlcBearerconfigR15SetupLogicalchannelidentityconfigR15ChoiceNothing = iota
	RlcBearerconfigR15SetupLogicalchannelidentityconfigR15ChoiceLogicalchannelidentityR15
	RlcBearerconfigR15SetupLogicalchannelidentityconfigR15ChoiceLogicalchannelidentityextR15
)

type RlcBearerconfigR15SetupLogicalchannelidentityconfigR15 struct {
	Choice                       uint64
	LogicalchannelidentityR15    *utils.INTEGER `lb:1,ub:10`
	LogicalchannelidentityextR15 *utils.INTEGER `lb:32,ub:38`
}
