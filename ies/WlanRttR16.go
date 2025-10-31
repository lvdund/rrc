package ies

import "rrc/utils"

// WLAN-RTT-r16 ::= SEQUENCE
// Extensible
type WlanRttR16 struct {
	RttvalueR16    utils.INTEGER `lb:0,ub:16777215`
	RttunitsR16    WlanRttR16RttunitsR16
	RttaccuracyR16 *utils.INTEGER `lb:0,ub:255`
}
