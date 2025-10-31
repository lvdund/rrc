package ies

import "rrc/utils"

// BeamMeasConfigIdle-NR-r16-reportQuantityRS-Indexes-r16 ::= ENUMERATED
type BeammeasconfigidleNrR16ReportquantityrsIndexesR16 struct {
	Value utils.ENUMERATED
}

const (
	BeammeasconfigidleNrR16ReportquantityrsIndexesR16EnumeratedNothing = iota
	BeammeasconfigidleNrR16ReportquantityrsIndexesR16EnumeratedRsrp
	BeammeasconfigidleNrR16ReportquantityrsIndexesR16EnumeratedRsrq
	BeammeasconfigidleNrR16ReportquantityrsIndexesR16EnumeratedBoth
)
