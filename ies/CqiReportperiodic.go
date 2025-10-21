package ies

import "rrc/utils"

// CQI-ReportPeriodic ::= CHOICE
type CqiReportperiodic interface {
	isCqiReportperiodic()
}

type CqiReportperiodicRelease struct {
	Value struct{}
}

func (*CqiReportperiodicRelease) isCqiReportperiodic() {}

type CqiReportperiodicSetup struct {
	Value interface{}
}

func (*CqiReportperiodicSetup) isCqiReportperiodic() {}
