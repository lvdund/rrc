package ies

import "rrc/utils"

// UECapabilityEnquiry-v1310-IEs ::= SEQUENCE
type UecapabilityenquiryV1310 struct {
	RequestreducedformatR13         *bool
	RequestskipfallbackcombR13      *bool
	RequestedmaxccsdlR13            *utils.INTEGER `lb:0,ub:32`
	RequestedmaxccsulR13            *utils.INTEGER `lb:0,ub:32`
	RequestreducedintnoncontcombR13 *bool
	Noncriticalextension            *UecapabilityenquiryV1430
}
