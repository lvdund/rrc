package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-increasedRepetition-r17 ::= ENUMERATED
type MimoParametersperbandSrsIncreasedrepetitionR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsIncreasedrepetitionR17EnumeratedNothing = iota
	MimoParametersperbandSrsIncreasedrepetitionR17EnumeratedSupported
)
