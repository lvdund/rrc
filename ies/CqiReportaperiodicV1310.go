package ies

import "rrc/utils"

// CQI-ReportAperiodic-v1310 ::= CHOICE
type CqiReportaperiodicV1310 interface {
	isCqiReportaperiodicV1310()
}

type CqiReportaperiodicV1310Release struct {
	Value struct{}
}

func (*CqiReportaperiodicV1310Release) isCqiReportaperiodicV1310() {}

type CqiReportaperiodicV1310Setup struct {
	Value interface{}
}

func (*CqiReportaperiodicV1310Setup) isCqiReportaperiodicV1310() {}
