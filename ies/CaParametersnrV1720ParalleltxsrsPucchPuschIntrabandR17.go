package ies

import "rrc/utils"

// CA-ParametersNR-v1720-parallelTxSRS-PUCCH-PUSCH-intraBand-r17 ::= ENUMERATED
type CaParametersnrV1720ParalleltxsrsPucchPuschIntrabandR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720ParalleltxsrsPucchPuschIntrabandR17EnumeratedNothing = iota
	CaParametersnrV1720ParalleltxsrsPucchPuschIntrabandR17EnumeratedSupported
)
