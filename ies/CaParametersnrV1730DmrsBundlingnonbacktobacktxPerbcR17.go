package ies

import "rrc/utils"

// CA-ParametersNR-v1730-dmrs-BundlingNonBackToBackTX-PerBC-r17 ::= ENUMERATED
type CaParametersnrV1730DmrsBundlingnonbacktobacktxPerbcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730DmrsBundlingnonbacktobacktxPerbcR17EnumeratedNothing = iota
	CaParametersnrV1730DmrsBundlingnonbacktobacktxPerbcR17EnumeratedSupported
)
