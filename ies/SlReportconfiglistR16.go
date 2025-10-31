package ies

// SL-ReportConfigList-r16 ::= SEQUENCE OF SL-ReportConfigInfo-r16
// SIZE (1..maxNrofSL-ReportConfigId-r16)
type SlReportconfiglistR16 struct {
	Value []SlReportconfiginfoR16 `lb:1,ub:maxNrofSLReportconfigidR16`
}
