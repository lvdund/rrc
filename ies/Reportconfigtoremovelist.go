package ies

// ReportConfigToRemoveList ::= SEQUENCE OF ReportConfigId
// SIZE (1..maxReportConfigId)
type Reportconfigtoremovelist struct {
	Value []Reportconfigid `lb:1,ub:maxReportConfigId`
}
