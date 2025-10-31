package ies

// CQI-ReportPeriodic-setup-cqi-FormatIndicatorPeriodic ::= CHOICE
const (
	CqiReportperiodicSetupCqiFormatindicatorperiodicChoiceNothing = iota
	CqiReportperiodicSetupCqiFormatindicatorperiodicChoiceWidebandcqi
	CqiReportperiodicSetupCqiFormatindicatorperiodicChoiceSubbandcqi
)

type CqiReportperiodicSetupCqiFormatindicatorperiodic struct {
	Choice      uint64
	Widebandcqi *struct{}
	Subbandcqi  *CqiReportperiodicSetupCqiFormatindicatorperiodicSubbandcqi
}
