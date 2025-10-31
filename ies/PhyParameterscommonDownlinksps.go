package ies

import "rrc/utils"

// Phy-ParametersCommon-downlinkSPS ::= ENUMERATED
type PhyParameterscommonDownlinksps struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDownlinkspsEnumeratedNothing = iota
	PhyParameterscommonDownlinkspsEnumeratedSupported
)
