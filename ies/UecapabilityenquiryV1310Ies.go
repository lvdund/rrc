package ies

import "rrc/utils"

// UECapabilityEnquiry-v1310-IEs ::= SEQUENCE
type UecapabilityenquiryV1310Ies struct {
	RequestreducedformatR13         *utils.ENUMERATED
	RequestskipfallbackcombR13      *utils.ENUMERATED
	RequestedmaxccsdlR13            *utils.INTEGER
	RequestedmaxccsulR13            *utils.INTEGER
	RequestreducedintnoncontcombR13 *utils.ENUMERATED
	Noncriticalextension            *UecapabilityenquiryV1430Ies
}
