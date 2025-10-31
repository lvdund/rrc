package ies

import "rrc/utils"

// MIMO-ParametersPerBand-semi-PersistentL1-SINR-Report-PUCCH-r16-supportReportFormat4-14OFDM-syms-r16 ::= ENUMERATED
type MimoParametersperbandSemiPersistentl1SinrReportPucchR16Supportreportformat414ofdmSymsR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSemiPersistentl1SinrReportPucchR16Supportreportformat414ofdmSymsR16EnumeratedNothing = iota
	MimoParametersperbandSemiPersistentl1SinrReportPucchR16Supportreportformat414ofdmSymsR16EnumeratedSupported
)
