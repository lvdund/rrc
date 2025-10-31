package ies

import "rrc/utils"

// Phy-ParametersCommon-sp-CSI-ReportPUCCH ::= ENUMERATED
type PhyParameterscommonSpCsiReportpucch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSpCsiReportpucchEnumeratedNothing = iota
	PhyParameterscommonSpCsiReportpucchEnumeratedSupported
)
