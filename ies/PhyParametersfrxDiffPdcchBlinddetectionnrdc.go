package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pdcch-BlindDetectionNRDC ::= SEQUENCE
type PhyParametersfrxDiffPdcchBlinddetectionnrdc struct {
	PdcchBlinddetectionmcgUe utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionscgUe utils.INTEGER `lb:0,ub:15`
}
