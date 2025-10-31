package ies

import "rrc/utils"

// CA-ParametersNR-v1610-pdcch-BlindDetectionMCG-UE-Mixed-r16 ::= SEQUENCE
type CaParametersnrV1610PdcchBlinddetectionmcgUeMixedR16 struct {
	PdcchBlinddetectionmcgUe1R16 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionmcgUe2R16 utils.INTEGER `lb:0,ub:15`
}
