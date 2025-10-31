package ies

import "rrc/utils"

// UE-CapabilityRAT-Request ::= SEQUENCE
// Extensible
type UeCapabilityratRequest struct {
	RatType                 RatType
	Capabilityrequestfilter *utils.OCTETSTRING
}
