package ies

// UE-CapabilityRequest ::= SEQUENCE OF RAT-Type
// SIZE (1..maxRAT-Capabilities)
type UeCapabilityrequest struct {
	Value []RatType `lb:1,ub:maxRATCapabilities`
}
