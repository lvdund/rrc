package ies

import "rrc/utils"

// WLAN-CarrierInfo-r13 ::= SEQUENCE
// Extensible
type WlanCarrierinfoR13 struct {
	OperatingclassR13 *utils.INTEGER
	CountrycodeR13    *utils.ENUMERATED
	ChannelnumbersR13 *WlanChannellistR13
}
