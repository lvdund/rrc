package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedSeparateTCI-commonMultiCC-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedseparatetciCommonmulticcR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedseparatetciCommonmulticcR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedseparatetciCommonmulticcR17EnumeratedSupported
)
