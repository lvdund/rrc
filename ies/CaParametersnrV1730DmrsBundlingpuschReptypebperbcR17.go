package ies

import "rrc/utils"

// CA-ParametersNR-v1730-dmrs-BundlingPUSCH-RepTypeBPerBC-r17 ::= ENUMERATED
type CaParametersnrV1730DmrsBundlingpuschReptypebperbcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730DmrsBundlingpuschReptypebperbcR17EnumeratedNothing = iota
	CaParametersnrV1730DmrsBundlingpuschReptypebperbcR17EnumeratedSupported
)
