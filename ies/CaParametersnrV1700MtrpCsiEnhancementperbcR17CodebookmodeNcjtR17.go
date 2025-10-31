package ies

import "rrc/utils"

// CA-ParametersNR-v1700-mTRP-CSI-EnhancementPerBC-r17-codebookMode-NCJT-r17 ::= ENUMERATED
type CaParametersnrV1700MtrpCsiEnhancementperbcR17CodebookmodeNcjtR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700MtrpCsiEnhancementperbcR17CodebookmodeNcjtR17EnumeratedNothing = iota
	CaParametersnrV1700MtrpCsiEnhancementperbcR17CodebookmodeNcjtR17EnumeratedMode1
	CaParametersnrV1700MtrpCsiEnhancementperbcR17CodebookmodeNcjtR17EnumeratedMode1and2
)
