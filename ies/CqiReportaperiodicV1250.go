package ies

import "rrc/utils"

// CQI-ReportAperiodic-v1250 ::= CHOICE
type CqiReportaperiodicV1250 interface {
	isCqiReportaperiodicV1250()
}

type CqiReportaperiodicV1250Release struct {
	Value struct{}
}

func (*CqiReportaperiodicV1250Release) isCqiReportaperiodicV1250() {}

type CqiReportaperiodicV1250Setup struct {
	Value interface{}
}

func (*CqiReportaperiodicV1250Setup) isCqiReportaperiodicV1250() {}
