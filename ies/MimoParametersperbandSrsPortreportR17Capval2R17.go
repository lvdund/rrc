package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-PortReport-r17-capVal2-r17 ::= ENUMERATED
type MimoParametersperbandSrsPortreportR17Capval2R17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsPortreportR17Capval2R17EnumeratedNothing = iota
	MimoParametersperbandSrsPortreportR17Capval2R17EnumeratedN1
	MimoParametersperbandSrsPortreportR17Capval2R17EnumeratedN2
	MimoParametersperbandSrsPortreportR17Capval2R17EnumeratedN4
)
