package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-PortReport-r17-capVal1-r17 ::= ENUMERATED
type MimoParametersperbandSrsPortreportR17Capval1R17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsPortreportR17Capval1R17EnumeratedNothing = iota
	MimoParametersperbandSrsPortreportR17Capval1R17EnumeratedN1
	MimoParametersperbandSrsPortreportR17Capval1R17EnumeratedN2
	MimoParametersperbandSrsPortreportR17Capval1R17EnumeratedN4
)
