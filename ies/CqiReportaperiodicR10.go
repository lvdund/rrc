package ies

import "rrc/utils"

// CQI-ReportAperiodic-r10 ::= CHOICE
type CqiReportaperiodicR10 interface {
	isCqiReportaperiodicR10()
}

type CqiReportaperiodicR10Release struct {
	Value struct{}
}

func (*CqiReportaperiodicR10Release) isCqiReportaperiodicR10() {}

type CqiReportaperiodicR10Setup struct {
	Value interface{}
}

func (*CqiReportaperiodicR10Setup) isCqiReportaperiodicR10() {}
