package ies

import "rrc/utils"

// CA-ParametersNR-v1720-parallelTxPRACH-SRS-PUCCH-PUSCH-intraBand-r17 ::= ENUMERATED
type CaParametersnrV1720ParalleltxprachSrsPucchPuschIntrabandR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720ParalleltxprachSrsPucchPuschIntrabandR17EnumeratedNothing = iota
	CaParametersnrV1720ParalleltxprachSrsPucchPuschIntrabandR17EnumeratedSupported
)
