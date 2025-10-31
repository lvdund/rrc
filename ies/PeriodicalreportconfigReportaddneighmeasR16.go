package ies

import "rrc/utils"

// PeriodicalReportConfig-reportAddNeighMeas-r16 ::= ENUMERATED
type PeriodicalreportconfigReportaddneighmeasR16 struct {
	Value utils.ENUMERATED
}

const (
	PeriodicalreportconfigReportaddneighmeasR16EnumeratedNothing = iota
	PeriodicalreportconfigReportaddneighmeasR16EnumeratedSetup
)
