package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-cli-RSSI-FDM-DL-r16 ::= ENUMERATED
type PhyParametersfrxDiffCliRssiFdmDlR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffCliRssiFdmDlR16EnumeratedNothing = iota
	PhyParametersfrxDiffCliRssiFdmDlR16EnumeratedSupported
)
