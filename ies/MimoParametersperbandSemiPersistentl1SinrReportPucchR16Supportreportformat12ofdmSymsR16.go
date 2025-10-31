package ies

import "rrc/utils"

// MIMO-ParametersPerBand-semi-PersistentL1-SINR-Report-PUCCH-r16-supportReportFormat1-2OFDM-syms-r16 ::= ENUMERATED
type MimoParametersperbandSemiPersistentl1SinrReportPucchR16Supportreportformat12ofdmSymsR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSemiPersistentl1SinrReportPucchR16Supportreportformat12ofdmSymsR16EnumeratedNothing = iota
	MimoParametersperbandSemiPersistentl1SinrReportPucchR16Supportreportformat12ofdmSymsR16EnumeratedSupported
)
