package ies

import "rrc/utils"

// PDCCH-BlindDetectionCG-UE-Mixed-r17 ::= SEQUENCE
type PdcchBlinddetectioncgUeMixedR17 struct {
	PdcchBlinddetectioncgUe1R17 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectioncgUe2R17 utils.INTEGER `lb:0,ub:15`
}
