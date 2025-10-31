package ies

import "rrc/utils"

// CA-ParametersNR-v1610 ::= SEQUENCE
type CaParametersnrV1610 struct {
	ParalleltxmsgaSrsPucchPuschR16          *CaParametersnrV1610ParalleltxmsgaSrsPucchPuschR16
	MsgaSulR16                              *CaParametersnrV1610MsgaSulR16
	JointsearchspaceswitchacrosscellsR16    *CaParametersnrV1610JointsearchspaceswitchacrosscellsR16
	HalfDuplextddCaSamescsR16               *CaParametersnrV1610HalfDuplextddCaSamescsR16
	ScelldormancywithinactivetimeR16        *CaParametersnrV1610ScelldormancywithinactivetimeR16
	ScelldormancyoutsideactivetimeR16       *CaParametersnrV1610ScelldormancyoutsideactivetimeR16
	CrosscarrieraCsiTrigdiffscsR16          *CaParametersnrV1610CrosscarrieraCsiTrigdiffscsR16
	DefaultqclCrosscarrieraCsiTrigR16       *CaParametersnrV1610DefaultqclCrosscarrieraCsiTrigR16
	IntercaNonalignedframeR16               *CaParametersnrV1610IntercaNonalignedframeR16
	SimulSrsTransBcR16                      *CaParametersnrV1610SimulSrsTransBcR16
	InterfreqdapsR16                        *CaParametersnrV1610InterfreqdapsR16
	CodebookparametersperbcR16              *CodebookparametersV1610
	BlinddetectfactorR16                    *utils.INTEGER `lb:0,ub:2`
	PdcchMonitoringcaR16                    *CaParametersnrV1610PdcchMonitoringcaR16
	PdcchBlinddetectioncaMixedR16           *CaParametersnrV1610PdcchBlinddetectioncaMixedR16
	PdcchBlinddetectionmcgUeR16             *utils.INTEGER `lb:0,ub:14`
	PdcchBlinddetectionscgUeR16             *utils.INTEGER `lb:0,ub:14`
	PdcchBlinddetectionmcgUeMixedR16        *CaParametersnrV1610PdcchBlinddetectionmcgUeMixedR16
	PdcchBlinddetectionscgUeMixedR16        *CaParametersnrV1610PdcchBlinddetectionscgUeMixedR16
	CrosscarrierschedulingdlDiffscsR16      *CaParametersnrV1610CrosscarrierschedulingdlDiffscsR16
	CrosscarrierschedulingdefaultqclR16     *CaParametersnrV1610CrosscarrierschedulingdefaultqclR16
	CrosscarrierschedulingulDiffscsR16      *CaParametersnrV1610CrosscarrierschedulingulDiffscsR16
	SimulSrsMimoTransBcR16                  *CaParametersnrV1610SimulSrsMimoTransBcR16
	CodebookparametersadditionperbcR16      *CodebookparametersadditionperbcR16
	CodebookcomboparametersadditionperbcR16 *CodebookcomboparametersadditionperbcR16
}
