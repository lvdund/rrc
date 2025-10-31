package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-csi-ReportWithoutCQI ::= ENUMERATED
type PhyParametersfrxDiffCsiReportwithoutcqi struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffCsiReportwithoutcqiEnumeratedNothing = iota
	PhyParametersfrxDiffCsiReportwithoutcqiEnumeratedSupported
)
