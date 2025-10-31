package ies

// UERadioPagingInformation-criticalExtensions ::= CHOICE
const (
	UeradiopaginginformationCriticalextensionsChoiceNothing = iota
	UeradiopaginginformationCriticalextensionsChoiceC1
	UeradiopaginginformationCriticalextensionsChoiceCriticalextensionsfuture
)

type UeradiopaginginformationCriticalextensions struct {
	Choice                   uint64
	C1                       *UeradiopaginginformationCriticalextensionsC1
	Criticalextensionsfuture *UeradiopaginginformationCriticalextensionsCriticalextensionsfuture
}
