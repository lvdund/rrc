package ies

import "rrc/utils"

// Phy-ParametersCommon-dynamicPRB-BundlingDL ::= ENUMERATED
type PhyParameterscommonDynamicprbBundlingdl struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDynamicprbBundlingdlEnumeratedNothing = iota
	PhyParameterscommonDynamicprbBundlingdlEnumeratedSupported
)
