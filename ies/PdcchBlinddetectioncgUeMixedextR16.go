package ies

import "rrc/utils"

// PDCCH-BlindDetectionCG-UE-MixedExt-r16 ::= SEQUENCE
type PdcchBlinddetectioncgUeMixedextR16 struct {
	PdcchBlinddetectioncgUe1R16 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectioncgUe2R16 utils.INTEGER `lb:0,ub:15`
}
