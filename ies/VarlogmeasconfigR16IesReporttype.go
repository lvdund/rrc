package ies

// VarLogMeasConfig-r16-IEs-reportType ::= CHOICE
const (
	VarlogmeasconfigR16IesReporttypeChoiceNothing = iota
	VarlogmeasconfigR16IesReporttypeChoicePeriodical
	VarlogmeasconfigR16IesReporttypeChoiceEventtriggered
)

type VarlogmeasconfigR16IesReporttype struct {
	Choice         uint64
	Periodical     *LoggedperiodicalreportconfigR16
	Eventtriggered *LoggedeventtriggerconfigR16
}
