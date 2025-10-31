package ies

// CSI-ReportConfig-reportConfigType-periodic ::= SEQUENCE
type CsiReportconfigReportconfigtypePeriodic struct {
	Reportslotconfig     CsiReportperiodicityandoffset
	PucchCsiResourcelist []PucchCsiResource `lb:1,ub:maxNrofBWPs`
}
