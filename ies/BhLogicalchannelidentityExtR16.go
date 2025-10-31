package ies

import "rrc/utils"

// BH-LogicalChannelIdentity-Ext-r16 ::= utils.INTEGER (320.. maxLC-ID-Iab-r16)
type BhLogicalchannelidentityExtR16 struct {
	Value utils.INTEGER `lb:0,ub:maxLCIdIabR16`
}
