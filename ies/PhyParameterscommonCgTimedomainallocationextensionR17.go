package ies

import "rrc/utils"

// Phy-ParametersCommon-cg-TimeDomainAllocationExtension-r17 ::= ENUMERATED
type PhyParameterscommonCgTimedomainallocationextensionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCgTimedomainallocationextensionR17EnumeratedNothing = iota
	PhyParameterscommonCgTimedomainallocationextensionR17EnumeratedSupported
)
