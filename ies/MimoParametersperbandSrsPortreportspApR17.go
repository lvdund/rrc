package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-PortReportSP-AP-r17 ::= ENUMERATED
type MimoParametersperbandSrsPortreportspApR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsPortreportspApR17EnumeratedNothing = iota
	MimoParametersperbandSrsPortreportspApR17EnumeratedSupported
)
