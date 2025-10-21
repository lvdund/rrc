package ies

import "rrc/utils"

// MeasIdToAddModExt-r12 ::= SEQUENCE
type MeasidtoaddmodextR12 struct {
	MeasidV1250       MeasidV1250
	MeasobjectidR12   Measobjectid
	ReportconfigidR12 Reportconfigid
}
