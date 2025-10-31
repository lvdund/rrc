package ies

import "rrc/utils"

// Phy-ParametersCommon-priorityIndicatorInDCI-Multicast-r17 ::= ENUMERATED
type PhyParameterscommonPriorityindicatorindciMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPriorityindicatorindciMulticastR17EnumeratedNothing = iota
	PhyParameterscommonPriorityindicatorindciMulticastR17EnumeratedSupported
)
