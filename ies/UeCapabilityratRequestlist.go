package ies

// UE-CapabilityRAT-RequestList ::= SEQUENCE OF UE-CapabilityRAT-Request
// SIZE (1..maxRAT-CapabilityContainers)
type UeCapabilityratRequestlist struct {
	Value []UeCapabilityratRequest `lb:1,ub:maxRATCapabilitycontainers`
}
