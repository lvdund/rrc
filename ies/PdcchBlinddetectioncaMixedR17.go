package ies

import "rrc/utils"

// PDCCH-BlindDetectionCA-Mixed-r17 ::= SEQUENCE
type PdcchBlinddetectioncaMixedR17 struct {
	PdcchBlinddetectionca1R17 *utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionca2R17 *utils.INTEGER `lb:0,ub:15`
}
