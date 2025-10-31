package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedSeparateTCI-perBWP-CA-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedseparatetciPerbwpCaR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedseparatetciPerbwpCaR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedseparatetciPerbwpCaR17EnumeratedSupported
)
