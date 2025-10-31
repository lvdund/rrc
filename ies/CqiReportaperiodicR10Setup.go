package ies

// CQI-ReportAperiodic-r10-setup ::= SEQUENCE
type CqiReportaperiodicR10Setup struct {
	CqiReportmodeaperiodicR10 CqiReportmodeaperiodic
	AperiodiccsiTriggerR10    *CqiReportaperiodicR10SetupAperiodiccsiTriggerR10
}
