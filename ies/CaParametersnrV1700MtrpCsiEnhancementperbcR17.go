package ies

import "rrc/utils"

// CA-ParametersNR-v1700-mTRP-CSI-EnhancementPerBC-r17 ::= SEQUENCE
type CaParametersnrV1700MtrpCsiEnhancementperbcR17 struct {
	MaxnumnzpCsiRsR17          utils.INTEGER `lb:0,ub:8`
	CsiReportModeR17           CaParametersnrV1700MtrpCsiEnhancementperbcR17CsiReportModeR17
	SupportedcomboacrossccsR17 []CsiMultitrpSupportedcombinationsR17 `lb:1,ub:16`
	CodebookmodeNcjtR17        CaParametersnrV1700MtrpCsiEnhancementperbcR17CodebookmodeNcjtR17
}
