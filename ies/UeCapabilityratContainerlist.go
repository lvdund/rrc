package ies

import "rrc/utils"

// UE-CapabilityRAT-ContainerList ::= SEQUENCE OF UE-CapabilityRAT-Container
// SIZE (0..maxRAT-Capabilities)
type UeCapabilityratContainerlist struct {
	Value utils.Sequence[UeCapabilityratContainer]
}
