package ies

// UECapabilityEnquiry-criticalExtensions-c1 ::= CHOICE
const (
	UecapabilityenquiryCriticalextensionsC1ChoiceNothing = iota
	UecapabilityenquiryCriticalextensionsC1ChoiceUecapabilityenquiryR8
	UecapabilityenquiryCriticalextensionsC1ChoiceSpare3
	UecapabilityenquiryCriticalextensionsC1ChoiceSpare2
	UecapabilityenquiryCriticalextensionsC1ChoiceSpare1
)

type UecapabilityenquiryCriticalextensionsC1 struct {
	Choice                uint64
	UecapabilityenquiryR8 *UecapabilityenquiryR8
	Spare3                *struct{}
	Spare2                *struct{}
	Spare1                *struct{}
}
