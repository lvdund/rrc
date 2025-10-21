package ies

import "rrc/utils"

// UECapabilityEnquiry-v1530-IEs ::= SEQUENCE
type UecapabilityenquiryV1530Ies struct {
	RequeststtiSptCapabilityR15 *utils.ENUMERATED
	EutraNrOnlyR15              *utils.ENUMERATED
	Noncriticalextension        *UecapabilityenquiryV1550Ies
}
