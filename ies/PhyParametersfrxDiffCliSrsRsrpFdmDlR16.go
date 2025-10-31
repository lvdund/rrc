package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-cli-SRS-RSRP-FDM-DL-r16 ::= ENUMERATED
type PhyParametersfrxDiffCliSrsRsrpFdmDlR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffCliSrsRsrpFdmDlR16EnumeratedNothing = iota
	PhyParametersfrxDiffCliSrsRsrpFdmDlR16EnumeratedSupported
)
