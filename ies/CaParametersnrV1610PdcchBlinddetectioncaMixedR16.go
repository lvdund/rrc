package ies

import "rrc/utils"

// CA-ParametersNR-v1610-pdcch-BlindDetectionCA-Mixed-r16 ::= SEQUENCE
type CaParametersnrV1610PdcchBlinddetectioncaMixedR16 struct {
	PdcchBlinddetectionca1R16   utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionca2R16   utils.INTEGER `lb:0,ub:15`
	SupportedspanarrangementR16 CaParametersnrV1610PdcchBlinddetectioncaMixedR16SupportedspanarrangementR16
}
