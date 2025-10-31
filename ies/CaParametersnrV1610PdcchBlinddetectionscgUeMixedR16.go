package ies

import "rrc/utils"

// CA-ParametersNR-v1610-pdcch-BlindDetectionSCG-UE-Mixed-r16 ::= SEQUENCE
type CaParametersnrV1610PdcchBlinddetectionscgUeMixedR16 struct {
	PdcchBlinddetectionscgUe1R16 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionscgUe2R16 utils.INTEGER `lb:0,ub:15`
}
