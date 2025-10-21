package ies

import "rrc/utils"

// UECapabilityEnquiry-v1610-IEs ::= SEQUENCE
type UecapabilityenquiryV1610Ies struct {
	RrcSegallowedR16     *utils.ENUMERATED
	Noncriticalextension *interface{}
}
