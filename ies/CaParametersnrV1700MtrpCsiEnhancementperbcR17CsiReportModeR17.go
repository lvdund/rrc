package ies

import "rrc/utils"

// CA-ParametersNR-v1700-mTRP-CSI-EnhancementPerBC-r17-cSI-Report-mode-r17 ::= ENUMERATED
type CaParametersnrV1700MtrpCsiEnhancementperbcR17CsiReportModeR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700MtrpCsiEnhancementperbcR17CsiReportModeR17EnumeratedNothing = iota
	CaParametersnrV1700MtrpCsiEnhancementperbcR17CsiReportModeR17EnumeratedMode1
	CaParametersnrV1700MtrpCsiEnhancementperbcR17CsiReportModeR17EnumeratedMode2
	CaParametersnrV1700MtrpCsiEnhancementperbcR17CsiReportModeR17EnumeratedBoth
)
