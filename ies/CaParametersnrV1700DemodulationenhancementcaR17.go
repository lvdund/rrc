package ies

import "rrc/utils"

// CA-ParametersNR-v1700-demodulationEnhancementCA-r17 ::= ENUMERATED
type CaParametersnrV1700DemodulationenhancementcaR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700DemodulationenhancementcaR17EnumeratedNothing = iota
	CaParametersnrV1700DemodulationenhancementcaR17EnumeratedSupported
)
