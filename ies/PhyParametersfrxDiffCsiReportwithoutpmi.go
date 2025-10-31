package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-csi-ReportWithoutPMI ::= ENUMERATED
type PhyParametersfrxDiffCsiReportwithoutpmi struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffCsiReportwithoutpmiEnumeratedNothing = iota
	PhyParametersfrxDiffCsiReportwithoutpmiEnumeratedSupported
)
