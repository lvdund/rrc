package ies

// UECapabilityEnquiry-r8-IEs ::= SEQUENCE
type UecapabilityenquiryR8 struct {
	UeCapabilityrequest  UeCapabilityrequest
	Noncriticalextension *UecapabilityenquiryV8a0
}
