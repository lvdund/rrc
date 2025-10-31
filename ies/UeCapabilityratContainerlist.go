package ies

// UE-CapabilityRAT-ContainerList ::= SEQUENCE OF UE-CapabilityRAT-Container
// SIZE (0..maxRAT-CapabilityContainers)
type UeCapabilityratContainerlist struct {
	Value []UeCapabilityratContainer `lb:0,ub:maxRATCapabilitycontainers`
}
