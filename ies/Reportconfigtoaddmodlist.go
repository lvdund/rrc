package ies

// ReportConfigToAddModList ::= SEQUENCE OF ReportConfigToAddMod
// SIZE (1..maxReportConfigId)
type Reportconfigtoaddmodlist struct {
	Value []Reportconfigtoaddmod `lb:1,ub:maxReportConfigId`
}
