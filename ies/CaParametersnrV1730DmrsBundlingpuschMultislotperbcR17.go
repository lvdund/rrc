package ies

import "rrc/utils"

// CA-ParametersNR-v1730-dmrs-BundlingPUSCH-multiSlotPerBC-r17 ::= ENUMERATED
type CaParametersnrV1730DmrsBundlingpuschMultislotperbcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730DmrsBundlingpuschMultislotperbcR17EnumeratedNothing = iota
	CaParametersnrV1730DmrsBundlingpuschMultislotperbcR17EnumeratedSupported
)
