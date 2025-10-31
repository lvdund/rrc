package ies

import "rrc/utils"

// MIMO-ParametersPerBand-pusch-TransCoherence ::= ENUMERATED
type MimoParametersperbandPuschTranscoherence struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandPuschTranscoherenceEnumeratedNothing = iota
	MimoParametersperbandPuschTranscoherenceEnumeratedNoncoherent
	MimoParametersperbandPuschTranscoherenceEnumeratedPartialcoherent
	MimoParametersperbandPuschTranscoherenceEnumeratedFullcoherent
)
