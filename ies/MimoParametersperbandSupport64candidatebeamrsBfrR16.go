package ies

import "rrc/utils"

// MIMO-ParametersPerBand-support64CandidateBeamRS-BFR-r16 ::= ENUMERATED
type MimoParametersperbandSupport64candidatebeamrsBfrR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSupport64candidatebeamrsBfrR16EnumeratedNothing = iota
	MimoParametersperbandSupport64candidatebeamrsBfrR16EnumeratedSupported
)
