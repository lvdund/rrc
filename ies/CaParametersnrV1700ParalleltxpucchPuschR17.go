package ies

import "rrc/utils"

// CA-ParametersNR-v1700-parallelTxPUCCH-PUSCH-r17 ::= ENUMERATED
type CaParametersnrV1700ParalleltxpucchPuschR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1700ParalleltxpucchPuschR17EnumeratedNothing = iota
	CaParametersnrV1700ParalleltxpucchPuschR17EnumeratedSupported
)
