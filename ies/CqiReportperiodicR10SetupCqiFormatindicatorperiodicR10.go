package ies

// CQI-ReportPeriodic-r10-setup-cqi-FormatIndicatorPeriodic-r10 ::= CHOICE
const (
	CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10ChoiceNothing = iota
	CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10ChoiceWidebandcqiR10
	CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10ChoiceSubbandcqiR10
)

type CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10 struct {
	Choice         uint64
	WidebandcqiR10 *CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10WidebandcqiR10
	SubbandcqiR10  *CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10SubbandcqiR10
}
