package ies

import "rrc/utils"

// CQI-ReportConfig-r15 ::= CHOICE
type CqiReportconfigR15 interface {
	isCqiReportconfigR15()
}

type CqiReportconfigR15Release struct {
	Value struct{}
}

func (*CqiReportconfigR15Release) isCqiReportconfigR15() {}

type CqiReportconfigR15Setup struct {
	Value interface{}
}

func (*CqiReportconfigR15Setup) isCqiReportconfigR15() {}
