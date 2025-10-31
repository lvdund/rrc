package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-startRB-locationHoppingPartial-r17 ::= ENUMERATED
type MimoParametersperbandSrsStartrbLocationhoppingpartialR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsStartrbLocationhoppingpartialR17EnumeratedNothing = iota
	MimoParametersperbandSrsStartrbLocationhoppingpartialR17EnumeratedSupported
)
