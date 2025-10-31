package ies

import "rrc/utils"

// Phy-ParametersCommon-sp-CSI-ReportPUSCH ::= ENUMERATED
type PhyParameterscommonSpCsiReportpusch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSpCsiReportpuschEnumeratedNothing = iota
	PhyParameterscommonSpCsiReportpuschEnumeratedSupported
)
