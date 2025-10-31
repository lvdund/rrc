package ies

import "rrc/utils"

// PDCCH-BlindDetectionMCG-SCG-r17 ::= SEQUENCE
type PdcchBlinddetectionmcgScgR17 struct {
	PdcchBlinddetectionmcgUeR17 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionscgUeR17 utils.INTEGER `lb:0,ub:15`
}
