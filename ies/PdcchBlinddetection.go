package ies

import "rrc/utils"

// PDCCH-BlindDetection ::= utils.INTEGER (1..15)
type PdcchBlinddetection struct {
	Value utils.INTEGER `lb:0,ub:15`
}
