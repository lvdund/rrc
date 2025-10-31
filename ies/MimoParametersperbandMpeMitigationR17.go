package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mpe-Mitigation-r17 ::= SEQUENCE
type MimoParametersperbandMpeMitigationR17 struct {
	MaxnumpMprRiPairsR17 utils.INTEGER `lb:0,ub:4`
	MaxnumconfrsR17      MimoParametersperbandMpeMitigationR17MaxnumconfrsR17
}
