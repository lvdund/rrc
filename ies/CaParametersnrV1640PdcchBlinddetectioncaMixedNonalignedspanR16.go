package ies

import "rrc/utils"

// CA-ParametersNR-v1640-pdcch-BlindDetectionCA-Mixed-NonAlignedSpan-r16 ::= SEQUENCE
type CaParametersnrV1640PdcchBlinddetectioncaMixedNonalignedspanR16 struct {
	PdcchBlinddetectionca1R16 utils.INTEGER `lb:0,ub:15`
	PdcchBlinddetectionca2R16 utils.INTEGER `lb:0,ub:15`
}
