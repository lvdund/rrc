package ies

// UERadioPagingInformation-NB-criticalExtensions ::= CHOICE
const (
	UeradiopaginginformationNbCriticalextensionsChoiceNothing = iota
	UeradiopaginginformationNbCriticalextensionsChoiceC1
	UeradiopaginginformationNbCriticalextensionsChoiceCriticalextensionsfuture
)

type UeradiopaginginformationNbCriticalextensions struct {
	Choice                   uint64
	C1                       *UeradiopaginginformationNbCriticalextensionsC1
	Criticalextensionsfuture *UeradiopaginginformationNbCriticalextensionsCriticalextensionsfuture
}
