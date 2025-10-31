package ies

import "rrc/utils"

// Phy-ParametersCommon-priorityIndicatorInDCI-SPS-Multicast-r17 ::= ENUMERATED
type PhyParameterscommonPriorityindicatorindciSpsMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPriorityindicatorindciSpsMulticastR17EnumeratedNothing = iota
	PhyParameterscommonPriorityindicatorindciSpsMulticastR17EnumeratedSupported
)
