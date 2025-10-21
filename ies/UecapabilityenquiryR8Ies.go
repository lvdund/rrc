package ies

import "rrc/utils"

// UECapabilityEnquiry-r8-IEs ::= SEQUENCE
type UecapabilityenquiryR8Ies struct {
	UeCapabilityrequest  UeCapabilityrequest
	Noncriticalextension *UecapabilityenquiryV8a0Ies
}
