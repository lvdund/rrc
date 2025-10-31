package ies

// CSI-ReportConfig-reportConfigType-semiPersistentOnPUCCH ::= SEQUENCE
type CsiReportconfigReportconfigtypeSemipersistentonpucch struct {
	Reportslotconfig     CsiReportperiodicityandoffset
	PucchCsiResourcelist []PucchCsiResource `lb:1,ub:maxNrofBWPs`
}
