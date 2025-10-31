package ies

// CQI-ReportConfig-r15-setup ::= SEQUENCE
type CqiReportconfigR15Setup struct {
	CqiReportconfigR10        *CqiReportconfigR10
	CqiReportconfigV1130      *CqiReportconfigV1130
	CqiReportconfigpcellV1250 *CqiReportconfigV1250
	CqiReportconfigV1310      *CqiReportconfigV1310
	CqiReportconfigV1320      *CqiReportconfigV1320
	CqiReportconfigV1430      *CqiReportconfigV1430
	AltcqiTable1024qamR15     *CqiReportconfigR15SetupAltcqiTable1024qamR15
}
