package ies

import "rrc/utils"

// CA-ParametersNR-v1730 ::= SEQUENCE
type CaParametersnrV1730 struct {
	DmrsBundlingpuschReptypeaperbcR17                  *CaParametersnrV1730DmrsBundlingpuschReptypeaperbcR17
	DmrsBundlingpuschReptypebperbcR17                  *CaParametersnrV1730DmrsBundlingpuschReptypebperbcR17
	DmrsBundlingpuschMultislotperbcR17                 *CaParametersnrV1730DmrsBundlingpuschMultislotperbcR17
	DmrsBundlingpucchRepperbcR17                       *CaParametersnrV1730DmrsBundlingpucchRepperbcR17
	DmrsBundlingrestartperbcR17                        *CaParametersnrV1730DmrsBundlingrestartperbcR17
	DmrsBundlingnonbacktobacktxPerbcR17                *CaParametersnrV1730DmrsBundlingnonbacktobacktxPerbcR17
	StayontargetccSrsCarrierswitchR17                  *CaParametersnrV1730StayontargetccSrsCarrierswitchR17
	FdmCodebookformuxUnicastmulticastharqAckR17        *CaParametersnrV1730FdmCodebookformuxUnicastmulticastharqAckR17
	Mode2TdmCodebookformuxUnicastmulticastharqAckR17   *CaParametersnrV1730Mode2TdmCodebookformuxUnicastmulticastharqAckR17
	Mode1Fortype1CodebookgenerationR17                 *CaParametersnrV1730Mode1Fortype1CodebookgenerationR17
	NackOnlyfeedbackspecificresourceforspsMulticastR17 *CaParametersnrV1730NackOnlyfeedbackspecificresourceforspsMulticastR17
	MultipucchConfigformulticastR17                    *CaParametersnrV1730MultipucchConfigformulticastR17
	PucchConfigforspsMulticastR17                      *CaParametersnrV1730PucchConfigforspsMulticastR17
	MaxnumbergRntiHarqAckCodebookR17                   *utils.INTEGER `lb:0,ub:4`
	MuxHarqAckUnicastmulticastR17                      *CaParametersnrV1730MuxHarqAckUnicastmulticastR17
}
