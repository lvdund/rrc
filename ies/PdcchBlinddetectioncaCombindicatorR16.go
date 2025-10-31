package ies

import "rrc/utils"

// PDCCH-BlindDetectionCA-CombIndicator-r16 ::= SEQUENCE
type PdcchBlinddetectioncaCombindicatorR16 struct {
	PdcchBlinddetectionca1R16 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionca2R16 utils.INTEGER `lb:0,ub:15`
}
