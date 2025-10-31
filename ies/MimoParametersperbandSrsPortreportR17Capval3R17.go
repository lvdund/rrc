package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-PortReport-r17-capVal3-r17 ::= ENUMERATED
type MimoParametersperbandSrsPortreportR17Capval3R17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsPortreportR17Capval3R17EnumeratedNothing = iota
	MimoParametersperbandSrsPortreportR17Capval3R17EnumeratedN1
	MimoParametersperbandSrsPortreportR17Capval3R17EnumeratedN2
	MimoParametersperbandSrsPortreportR17Capval3R17EnumeratedN4
)
