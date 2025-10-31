package ies

import "rrc/utils"

// CA-ParametersNR-v1720-parallelTxMsgA-SRS-PUCCH-PUSCH-intraBand-r17 ::= ENUMERATED
type CaParametersnrV1720ParalleltxmsgaSrsPucchPuschIntrabandR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720ParalleltxmsgaSrsPucchPuschIntrabandR17EnumeratedNothing = iota
	CaParametersnrV1720ParalleltxmsgaSrsPucchPuschIntrabandR17EnumeratedSupported
)
