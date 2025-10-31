package ies

import "rrc/utils"

// WLAN-RSSI-Range-r13 ::= utils.INTEGER (0..141)
type WlanRssiRangeR13 struct {
	Value utils.INTEGER `lb:0,ub:141`
}
