package ies

// ReportConfigNR-reportType ::= CHOICE
// Extensible
const (
	ReportconfignrReporttypeChoiceNothing = iota
	ReportconfignrReporttypeChoicePeriodical
	ReportconfignrReporttypeChoiceEventtriggered
	ReportconfignrReporttypeChoiceReportcgi
	ReportconfignrReporttypeChoiceReportsftd
	ReportconfignrReporttypeChoiceCondtriggerconfigR16
	ReportconfignrReporttypeChoiceCliPeriodicalR16
	ReportconfignrReporttypeChoiceCliEventtriggeredR16
	ReportconfignrReporttypeChoiceRxtxperiodicalR17
)

type ReportconfignrReporttype struct {
	Choice               uint64
	Periodical           *Periodicalreportconfig
	Eventtriggered       *Eventtriggerconfig
	Reportcgi            *Reportcgi
	Reportsftd           *ReportsftdNr
	CondtriggerconfigR16 *CondtriggerconfigR16
	CliPeriodicalR16     *CliPeriodicalreportconfigR16
	CliEventtriggeredR16 *CliEventtriggerconfigR16
	RxtxperiodicalR17    *RxtxperiodicalR17
}
