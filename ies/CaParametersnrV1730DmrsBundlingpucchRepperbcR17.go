package ies

import "rrc/utils"

// CA-ParametersNR-v1730-dmrs-BundlingPUCCH-RepPerBC-r17 ::= ENUMERATED
type CaParametersnrV1730DmrsBundlingpucchRepperbcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730DmrsBundlingpucchRepperbcR17EnumeratedNothing = iota
	CaParametersnrV1730DmrsBundlingpucchRepperbcR17EnumeratedSupported
)
