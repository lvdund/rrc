package ies

import "rrc/utils"

// MAC-ParametersCommon-enhancedUuDRX-forSidelink-r17 ::= ENUMERATED
type MacParameterscommonEnhanceduudrxForsidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonEnhanceduudrxForsidelinkR17EnumeratedNothing = iota
	MacParameterscommonEnhanceduudrxForsidelinkR17EnumeratedSupported
)
