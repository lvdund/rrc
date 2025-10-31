package ies

import "rrc/utils"

// WLAN-RSSI-Range-r16 ::= utils.INTEGER (0..141)
type WlanRssiRangeR16 struct {
	Value utils.INTEGER `lb:0,ub:141`
}
