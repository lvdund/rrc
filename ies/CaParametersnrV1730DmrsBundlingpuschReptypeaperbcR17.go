package ies

import "rrc/utils"

// CA-ParametersNR-v1730-dmrs-BundlingPUSCH-RepTypeAPerBC-r17 ::= ENUMERATED
type CaParametersnrV1730DmrsBundlingpuschReptypeaperbcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730DmrsBundlingpuschReptypeaperbcR17EnumeratedNothing = iota
	CaParametersnrV1730DmrsBundlingpuschReptypeaperbcR17EnumeratedSupported
)
