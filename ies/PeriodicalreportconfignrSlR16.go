package ies

// PeriodicalReportConfigNR-SL-r16 ::= SEQUENCE
// Extensible
type PeriodicalreportconfignrSlR16 struct {
	ReportintervalR16 Reportinterval
	ReportamountR16   PeriodicalreportconfignrSlR16ReportamountR16
	ReportquantityR16 MeasreportquantityR16
}
