package ies

import "rrc/utils"

// WLAN-Channel-r13 ::= utils.INTEGER (0..255)
type WlanChannelR13 struct {
	Value utils.INTEGER `lb:0,ub:255`
}
