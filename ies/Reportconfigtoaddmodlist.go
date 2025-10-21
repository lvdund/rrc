package ies

import "rrc/utils"

// ReportConfigToAddModList ::= SEQUENCE OF ReportConfigToAddMod
// SIZE (1..maxReportConfigId)
type Reportconfigtoaddmodlist struct {
	Value []Reportconfigtoaddmod
}
