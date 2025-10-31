package ies

import "rrc/utils"

// CA-ParametersNR-v1610-parallelTxMsgA-SRS-PUCCH-PUSCH-r16 ::= ENUMERATED
type CaParametersnrV1610ParalleltxmsgaSrsPucchPuschR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610ParalleltxmsgaSrsPucchPuschR16EnumeratedNothing = iota
	CaParametersnrV1610ParalleltxmsgaSrsPucchPuschR16EnumeratedSupported
)
