package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-partialFrequencySounding-r17 ::= ENUMERATED
type MimoParametersperbandSrsPartialfrequencysoundingR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsPartialfrequencysoundingR17EnumeratedNothing = iota
	MimoParametersperbandSrsPartialfrequencysoundingR17EnumeratedSupported
)
