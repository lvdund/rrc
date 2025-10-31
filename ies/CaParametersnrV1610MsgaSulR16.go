package ies

import "rrc/utils"

// CA-ParametersNR-v1610-msgA-SUL-r16 ::= ENUMERATED
type CaParametersnrV1610MsgaSulR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610MsgaSulR16EnumeratedNothing = iota
	CaParametersnrV1610MsgaSulR16EnumeratedSupported
)
