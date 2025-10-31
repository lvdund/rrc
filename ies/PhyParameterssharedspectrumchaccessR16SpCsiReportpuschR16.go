package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-sp-CSI-ReportPUSCH-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16SpCsiReportpuschR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16SpCsiReportpuschR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16SpCsiReportpuschR16EnumeratedSupported
)
