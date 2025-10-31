package ies

import "rrc/utils"

// RSSI-Range-r13 ::= utils.INTEGER (0..76)
type RssiRangeR13 struct {
	Value utils.INTEGER `lb:0,ub:76`
}
