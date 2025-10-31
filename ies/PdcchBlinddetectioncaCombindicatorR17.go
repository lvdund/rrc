package ies

import "rrc/utils"

// PDCCH-BlindDetectionCA-CombIndicator-r17 ::= SEQUENCE
type PdcchBlinddetectioncaCombindicatorR17 struct {
	PdcchBlinddetectionca1R17 *utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionca2R17 *utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionca3R17 utils.INTEGER  `lb:0,ub:15`
}
