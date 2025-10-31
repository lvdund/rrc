package ies

import "rrc/utils"

// MIMO-ParametersPerBand-multiDCI-multiTRP-Parameters-r16-outOfOrderOperationUL-r16 ::= ENUMERATED
type MimoParametersperbandMultidciMultitrpParametersR16OutoforderoperationulR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMultidciMultitrpParametersR16OutoforderoperationulR16EnumeratedNothing = iota
	MimoParametersperbandMultidciMultitrpParametersR16OutoforderoperationulR16EnumeratedSupported
)
