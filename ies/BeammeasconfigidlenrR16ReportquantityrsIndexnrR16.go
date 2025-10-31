package ies

import "rrc/utils"

// BeamMeasConfigIdleNR-r16-reportQuantityRS-IndexNR-r16 ::= ENUMERATED
type BeammeasconfigidlenrR16ReportquantityrsIndexnrR16 struct {
	Value utils.ENUMERATED
}

const (
	BeammeasconfigidlenrR16ReportquantityrsIndexnrR16EnumeratedNothing = iota
	BeammeasconfigidlenrR16ReportquantityrsIndexnrR16EnumeratedRsrp
	BeammeasconfigidlenrR16ReportquantityrsIndexnrR16EnumeratedRsrq
	BeammeasconfigidlenrR16ReportquantityrsIndexnrR16EnumeratedBoth
)
