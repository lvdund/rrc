package ies

import "rrc/utils"

// CRI-ReportConfig-r13 ::= CHOICE
type CriReportconfigR13 interface {
	isCriReportconfigR13()
}

type CriReportconfigR13Release struct {
	Value struct{}
}

func (*CriReportconfigR13Release) isCriReportconfigR13() {}

type CriReportconfigR13Setup struct {
	Value interface{}
}

func (*CriReportconfigR13Setup) isCriReportconfigR13() {}
