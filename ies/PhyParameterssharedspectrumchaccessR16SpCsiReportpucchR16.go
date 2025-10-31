package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-sp-CSI-ReportPUCCH-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16SpCsiReportpucchR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16SpCsiReportpucchR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16SpCsiReportpucchR16EnumeratedSupported
)
