package ies

import "rrc/utils"

// CQI-ReportPeriodic-r10 ::= CHOICE
type CqiReportperiodicR10 interface {
	isCqiReportperiodicR10()
}

type CqiReportperiodicR10Release struct {
	Value struct{}
}

func (*CqiReportperiodicR10Release) isCqiReportperiodicR10() {}

type CqiReportperiodicR10Setup struct {
	Value interface{}
}

func (*CqiReportperiodicR10Setup) isCqiReportperiodicR10() {}
