package ies

// CSI-ReportConfig-reportFreqConfiguration ::= SEQUENCE
// Extensible
type CsiReportconfigReportfreqconfiguration struct {
	CqiFormatindicator *CsiReportconfigReportfreqconfigurationCqiFormatindicator
	PmiFormatindicator *CsiReportconfigReportfreqconfigurationPmiFormatindicator
	CsiReportingband   *CsiReportconfigReportfreqconfigurationCsiReportingband
}
