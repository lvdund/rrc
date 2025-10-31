package ies

import "rrc/utils"

// MIMO-ParametersPerBand-multiDCI-multiTRP-Parameters-r16-separateCRS-RateMatching-r16 ::= ENUMERATED
type MimoParametersperbandMultidciMultitrpParametersR16SeparatecrsRatematchingR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMultidciMultitrpParametersR16SeparatecrsRatematchingR16EnumeratedNothing = iota
	MimoParametersperbandMultidciMultitrpParametersR16SeparatecrsRatematchingR16EnumeratedSupported
)
