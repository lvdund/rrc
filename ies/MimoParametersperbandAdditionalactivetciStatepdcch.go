package ies

import "rrc/utils"

// MIMO-ParametersPerBand-additionalActiveTCI-StatePDCCH ::= ENUMERATED
type MimoParametersperbandAdditionalactivetciStatepdcch struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandAdditionalactivetciStatepdcchEnumeratedNothing = iota
	MimoParametersperbandAdditionalactivetciStatepdcchEnumeratedSupported
)
