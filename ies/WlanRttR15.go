package ies

import "rrc/utils"

// WLAN-RTT-r15 ::= SEQUENCE
// Extensible
type WlanRttR15 struct {
	RttvalueR15    utils.INTEGER `lb:0,ub:16777215`
	RttunitsR15    WlanRttR15RttunitsR15
	RttaccuracyR15 *utils.INTEGER `lb:0,ub:255`
}
