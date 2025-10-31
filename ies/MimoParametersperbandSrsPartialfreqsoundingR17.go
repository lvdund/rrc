package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-partialFreqSounding-r17 ::= ENUMERATED
type MimoParametersperbandSrsPartialfreqsoundingR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsPartialfreqsoundingR17EnumeratedNothing = iota
	MimoParametersperbandSrsPartialfreqsoundingR17EnumeratedSupported
)
