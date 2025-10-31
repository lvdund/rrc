package ies

import "rrc/utils"

// IRAT-ParametersNR-v1610-nr-HO-ToEN-DC-r16 ::= ENUMERATED
type IratParametersnrV1610NrHoToenDcR16 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1610NrHoToenDcR16EnumeratedNothing = iota
	IratParametersnrV1610NrHoToenDcR16EnumeratedSupported
)
