package ies

// SL-ReportConfigToRemoveList-r16 ::= SEQUENCE OF SL-ReportConfigId-r16
// SIZE (1..maxNrofSL-ReportConfigId-r16)
type SlReportconfigtoremovelistR16 struct {
	Value []SlReportconfigidR16 `lb:1,ub:maxNrofSLReportconfigidR16`
}
