package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-PortReport-r17-capVal4-r17 ::= ENUMERATED
type MimoParametersperbandSrsPortreportR17Capval4R17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsPortreportR17Capval4R17EnumeratedNothing = iota
	MimoParametersperbandSrsPortreportR17Capval4R17EnumeratedN1
	MimoParametersperbandSrsPortreportR17Capval4R17EnumeratedN2
	MimoParametersperbandSrsPortreportR17Capval4R17EnumeratedN4
)
