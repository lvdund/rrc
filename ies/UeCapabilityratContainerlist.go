package ies

// UE-CapabilityRAT-ContainerList ::= SEQUENCE OF UE-CapabilityRAT-Container
// SIZE (0..maxRAT-Capabilities)
type UeCapabilityratContainerlist struct {
	Value []UeCapabilityratContainer `lb:0,ub:maxRATCapabilities`
}
