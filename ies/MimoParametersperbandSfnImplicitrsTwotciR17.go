package ies

import "rrc/utils"

// MIMO-ParametersPerBand-sfn-ImplicitRS-twoTCI-r17 ::= ENUMERATED
type MimoParametersperbandSfnImplicitrsTwotciR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSfnImplicitrsTwotciR17EnumeratedNothing = iota
	MimoParametersperbandSfnImplicitrsTwotciR17EnumeratedSupported
)
