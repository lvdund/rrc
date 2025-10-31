package ies

import "rrc/utils"

// CA-ParametersNR-v1730-dmrs-BundlingRestartPerBC-r17 ::= ENUMERATED
type CaParametersnrV1730DmrsBundlingrestartperbcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730DmrsBundlingrestartperbcR17EnumeratedNothing = iota
	CaParametersnrV1730DmrsBundlingrestartperbcR17EnumeratedSupported
)
