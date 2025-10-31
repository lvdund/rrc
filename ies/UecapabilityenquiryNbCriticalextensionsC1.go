package ies

// UECapabilityEnquiry-NB-criticalExtensions-c1 ::= CHOICE
const (
	UecapabilityenquiryNbCriticalextensionsC1ChoiceNothing = iota
	UecapabilityenquiryNbCriticalextensionsC1ChoiceUecapabilityenquiryR13
	UecapabilityenquiryNbCriticalextensionsC1ChoiceSpare1
)

type UecapabilityenquiryNbCriticalextensionsC1 struct {
	Choice                 uint64
	UecapabilityenquiryR13 *UecapabilityenquiryNbR13
	Spare1                 *struct{}
}
