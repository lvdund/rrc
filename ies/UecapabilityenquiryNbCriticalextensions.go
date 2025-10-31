package ies

// UECapabilityEnquiry-NB-criticalExtensions ::= CHOICE
const (
	UecapabilityenquiryNbCriticalextensionsChoiceNothing = iota
	UecapabilityenquiryNbCriticalextensionsChoiceC1
	UecapabilityenquiryNbCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityenquiryNbCriticalextensions struct {
	Choice                   uint64
	C1                       *UecapabilityenquiryNbCriticalextensionsC1
	Criticalextensionsfuture *UecapabilityenquiryNbCriticalextensionsCriticalextensionsfuture
}
