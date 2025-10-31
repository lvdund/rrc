package ies

import "rrc/utils"

// MIMO-ParametersPerBand-singleDCI-SDM-scheme-Parameters-r16-supportTwoPortDL-PTRS-r16 ::= ENUMERATED
type MimoParametersperbandSingledciSdmSchemeParametersR16SupporttwoportdlPtrsR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSingledciSdmSchemeParametersR16SupporttwoportdlPtrsR16EnumeratedNothing = iota
	MimoParametersperbandSingledciSdmSchemeParametersR16SupporttwoportdlPtrsR16EnumeratedSupported
)
