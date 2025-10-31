package ies

// LoggedMeasurementConfiguration-r16-IEs-reportType ::= CHOICE
// Extensible
const (
	LoggedmeasurementconfigurationR16IesReporttypeChoiceNothing = iota
	LoggedmeasurementconfigurationR16IesReporttypeChoicePeriodical
	LoggedmeasurementconfigurationR16IesReporttypeChoiceEventtriggered
)

type LoggedmeasurementconfigurationR16IesReporttype struct {
	Choice         uint64
	Periodical     *LoggedperiodicalreportconfigR16
	Eventtriggered *LoggedeventtriggerconfigR16
}
