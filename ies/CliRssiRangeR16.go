package ies

import "rrc/utils"

// CLI-RSSI-Range-r16 ::= utils.INTEGER (0..76)
type CliRssiRangeR16 struct {
	Value utils.INTEGER `lb:0,ub:76`
}
