package ies

import "rrc/utils"

// UE-CapabilityRequest ::= SEQUENCE OF RAT-Type
// SIZE (1..maxRAT-Capabilities)
type UeCapabilityrequest struct {
	Value []RatType
}
