package ies

// CQI-ReportConfigSCell-r15 ::= SEQUENCE
type CqiReportconfigscellR15 struct {
	CqiReportperiodicscellR15 *CqiReportperiodicscellR15
	AltcqiTable1024qamR15     *CqiReportconfigscellR15AltcqiTable1024qamR15
}
