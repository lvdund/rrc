package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-twoPHR-Reporting-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPuschTwophrReportingR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPuschTwophrReportingR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPuschTwophrReportingR17EnumeratedSupported
)
