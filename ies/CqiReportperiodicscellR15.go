package ies

import "rrc/utils"

// CQI-ReportPeriodicSCell-r15 ::= CHOICE
type CqiReportperiodicscellR15 interface {
	isCqiReportperiodicscellR15()
}

type CqiReportperiodicscellR15Release struct {
	Value struct{}
}

func (*CqiReportperiodicscellR15Release) isCqiReportperiodicscellR15() {}

type CqiReportperiodicscellR15Setup struct {
	Value interface{}
}

func (*CqiReportperiodicscellR15Setup) isCqiReportperiodicscellR15() {}
