package ies

import "rrc/utils"

// PDCCH-BlindDetectionCG-UE-Mixed1-r17 ::= SEQUENCE
type PdcchBlinddetectioncgUeMixed1R17 struct {
	PdcchBlinddetectioncgUe1R17 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectioncgUe2R17 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectioncgUe3R17 utils.INTEGER `lb:0,ub:15`
}
