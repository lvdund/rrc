package ies

// UECapabilityEnquiry-v1180-IEs ::= SEQUENCE
type UecapabilityenquiryV1180 struct {
	RequestedfrequencybandsR11 *[]FreqbandindicatorR11 `lb:1,ub:16`
	Noncriticalextension       *UecapabilityenquiryV1310
}
