package ies

import "rrc/utils"

// CA-ParametersNR-v1720 ::= SEQUENCE
type CaParametersnrV1720 struct {
	ParalleltxsrsPucchPuschIntrabandR17             *CaParametersnrV1720ParalleltxsrsPucchPuschIntrabandR17
	ParalleltxprachSrsPucchPuschIntrabandR17        *CaParametersnrV1720ParalleltxprachSrsPucchPuschIntrabandR17
	SemistaticpucchCellswitchsinglegroupR17         *CaParametersnrV1720SemistaticpucchCellswitchsinglegroupR17
	SemistaticpucchCellswitchtwogroupsR17           *[]TwopucchGrpConfigurationsR17 `lb:1,ub:maxTwoPUCCHGrpConfiglistR17`
	DynamicpucchCellswitchsamelengthsinglegroupR17  *CaParametersnrV1720DynamicpucchCellswitchsamelengthsinglegroupR17
	DynamicpucchCellswitchdifflengthsinglegroupR17  *CaParametersnrV1720DynamicpucchCellswitchdifflengthsinglegroupR17
	DynamicpucchCellswitchsamelengthtwogroupsR17    *[]TwopucchGrpConfigurationsR17 `lb:1,ub:maxTwoPUCCHGrpConfiglistR17`
	DynamicpucchCellswitchdifflengthtwogroupsR17    *[]TwopucchGrpConfigurationsR17 `lb:1,ub:maxTwoPUCCHGrpConfiglistR17`
	AckNackFeedbackformulticastR17                  *CaParametersnrV1720AckNackFeedbackformulticastR17
	PtpRetxMulticastR17                             *CaParametersnrV1720PtpRetxMulticastR17
	NackOnlyfeedbackformulticastR17                 *CaParametersnrV1720NackOnlyfeedbackformulticastR17
	NackOnlyfeedbackspecificresourceformulticastR17 *CaParametersnrV1720NackOnlyfeedbackspecificresourceformulticastR17
	AckNackFeedbackforspsMulticastR17               *CaParametersnrV1720AckNackFeedbackforspsMulticastR17
	PtpRetxSpsMulticastR17                          *CaParametersnrV1720PtpRetxSpsMulticastR17
	HigherpowerlimitR17                             *CaParametersnrV1720HigherpowerlimitR17
	ParalleltxmsgaSrsPucchPuschIntrabandR17         *CaParametersnrV1720ParalleltxmsgaSrsPucchPuschIntrabandR17
	PdcchMonitoringcaR17                            *utils.INTEGER                  `lb:0,ub:16`
	PdcchBlinddetectionmcgScgListR17                *[]PdcchBlinddetectionmcgScgR17 `lb:1,ub:maxNrofPdcchBlinddetectionR17`
	PdcchBlinddetectionmixedlist1R17                *[]PdcchBlinddetectionmixedR17  `lb:1,ub:maxNrofPdcchBlinddetectionR17`
	PdcchBlinddetectionmixedlist2R17                *[]PdcchBlinddetectionmixedR17  `lb:1,ub:maxNrofPdcchBlinddetectionR17`
	PdcchBlinddetectionmixedlist3R17                *[]PdcchBlinddetectionmixed1R17 `lb:1,ub:maxNrofPdcchBlinddetectionR17`
}
