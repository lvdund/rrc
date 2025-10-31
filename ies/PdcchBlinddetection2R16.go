package ies

import "rrc/utils"

// PDCCH-BlindDetection2-r16 ::= utils.INTEGER (1..15)
type PdcchBlinddetection2R16 struct {
	Value utils.INTEGER `lb:0,ub:15`
}
