package ies

// UECapabilityEnquiry-v1530-IEs ::= SEQUENCE
type UecapabilityenquiryV1530 struct {
	RequeststtiSptCapabilityR15 *bool
	EutraNrOnlyR15              *bool
	Noncriticalextension        *UecapabilityenquiryV1550
}
