package ies

import "rrc/utils"

// WLAN-CarrierInfo-r13 ::= SEQUENCE
// Extensible
type WlanCarrierinfoR13 struct {
	OperatingclassR13 *utils.INTEGER `lb:0,ub:255`
	CountrycodeR13    *WlanCarrierinfoR13CountrycodeR13
	ChannelnumbersR13 *WlanChannellistR13
}
