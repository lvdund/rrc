package ies

import "rrc/utils"

// PDCCH-BlindDetectionCA-MixedExt-r16 ::= SEQUENCE
type PdcchBlinddetectioncaMixedextR16 struct {
	PdcchBlinddetectionca1R16 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionca2R16 utils.INTEGER `lb:0,ub:15`
}
