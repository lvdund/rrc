package ies

import "rrc/utils"

// MIMO-ParametersPerBand-semi-PersistentL1-SINR-Report-PUSCH-r16 ::= ENUMERATED
type MimoParametersperbandSemiPersistentl1SinrReportPuschR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSemiPersistentl1SinrReportPuschR16EnumeratedNothing = iota
	MimoParametersperbandSemiPersistentl1SinrReportPuschR16EnumeratedSupported
)
