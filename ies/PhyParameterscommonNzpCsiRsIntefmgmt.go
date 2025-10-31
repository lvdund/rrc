package ies

import "rrc/utils"

// Phy-ParametersCommon-nzp-CSI-RS-IntefMgmt ::= ENUMERATED
type PhyParameterscommonNzpCsiRsIntefmgmt struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonNzpCsiRsIntefmgmtEnumeratedNothing = iota
	PhyParameterscommonNzpCsiRsIntefmgmtEnumeratedSupported
)
