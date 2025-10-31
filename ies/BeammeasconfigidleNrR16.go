package ies

import "rrc/utils"

// BeamMeasConfigIdle-NR-r16 ::= SEQUENCE
type BeammeasconfigidleNrR16 struct {
	ReportquantityrsIndexesR16  BeammeasconfigidleNrR16ReportquantityrsIndexesR16
	MaxnrofrsIndexestoreportR16 utils.INTEGER `lb:0,ub:maxNrofIndexesToReport`
	IncludebeammeasurementsR16  utils.BOOLEAN
}
