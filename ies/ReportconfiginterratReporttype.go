package ies

// ReportConfigInterRAT-reportType ::= CHOICE
// Extensible
const (
	ReportconfiginterratReporttypeChoiceNothing = iota
	ReportconfiginterratReporttypeChoicePeriodical
	ReportconfiginterratReporttypeChoiceEventtriggered
	ReportconfiginterratReporttypeChoiceReportcgi
	ReportconfiginterratReporttypeChoiceReportsftd
)

type ReportconfiginterratReporttype struct {
	Choice         uint64
	Periodical     *Periodicalreportconfiginterrat
	Eventtriggered *Eventtriggerconfiginterrat
	Reportcgi      *ReportcgiEutra
	Reportsftd     *ReportsftdEutra
}
