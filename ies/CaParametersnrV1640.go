package ies

import "rrc/utils"

// CA-ParametersNR-v1640 ::= SEQUENCE
type CaParametersnrV1640 struct {
	UplinktxdcTwocarrierreportR16                           *CaParametersnrV1640UplinktxdcTwocarrierreportR16
	Maxupto3diffNumerologiesconfigsinglepucchGrpR16         *PucchGrpCarriertypesR16
	Maxupto4diffNumerologiesconfigsinglepucchGrpR16         *PucchGrpCarriertypesR16
	TwopucchGrpConfigurationslistR16                        *[]TwopucchGrpConfigurationsR16 `lb:1,ub:maxTwoPUCCHGrpConfiglistR16`
	DiffnumerologyacrosspucchGroupCarriertypesR16           *CaParametersnrV1640DiffnumerologyacrosspucchGroupCarriertypesR16
	DiffnumerologywithinpucchGroupsmallerscsCarriertypesR16 *CaParametersnrV1640DiffnumerologywithinpucchGroupsmallerscsCarriertypesR16
	DiffnumerologywithinpucchGrouplargerscsCarriertypesR16  *CaParametersnrV1640DiffnumerologywithinpucchGrouplargerscsCarriertypesR16
	PdcchMonitoringcaNonalignedspanR16                      *utils.INTEGER `lb:0,ub:16`
	PdcchBlinddetectioncaMixedNonalignedspanR16             *CaParametersnrV1640PdcchBlinddetectioncaMixedNonalignedspanR16
}
