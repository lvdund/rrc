package ies

import "rrc/utils"

// UE-CapabilityRAT-Container ::= SEQUENCE
type UeCapabilityratContainer struct {
	RatType                  RatType
	UecapabilityratContainer utils.OCTETSTRING
}
