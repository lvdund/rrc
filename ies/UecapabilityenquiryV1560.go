package ies

// UECapabilityEnquiry-v1560-IEs ::= SEQUENCE
type UecapabilityenquiryV1560 struct {
	Capabilityrequestfiltercommon *UeCapabilityrequestfiltercommon
	Noncriticalextension          *UecapabilityenquiryV1610
}
