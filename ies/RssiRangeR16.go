package ies

import "rrc/utils"

// RSSI-Range-r16 ::= utils.INTEGER (0..76)
type RssiRangeR16 struct {
	Value utils.INTEGER `lb:0,ub:76`
}
