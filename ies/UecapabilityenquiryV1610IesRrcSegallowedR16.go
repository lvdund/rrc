package ies

import "rrc/utils"

// UECapabilityEnquiry-v1610-IEs-rrc-SegAllowed-r16 ::= ENUMERATED
type UecapabilityenquiryV1610IesRrcSegallowedR16 struct {
	Value utils.ENUMERATED
}

const (
	UecapabilityenquiryV1610IesRrcSegallowedR16EnumeratedNothing = iota
	UecapabilityenquiryV1610IesRrcSegallowedR16EnumeratedEnabled
)
