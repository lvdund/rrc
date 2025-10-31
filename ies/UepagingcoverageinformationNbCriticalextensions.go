package ies

// UEPagingCoverageInformation-NB-criticalExtensions ::= CHOICE
const (
	UepagingcoverageinformationNbCriticalextensionsChoiceNothing = iota
	UepagingcoverageinformationNbCriticalextensionsChoiceC1
	UepagingcoverageinformationNbCriticalextensionsChoiceCriticalextensionsfuture
)

type UepagingcoverageinformationNbCriticalextensions struct {
	Choice                   uint64
	C1                       *UepagingcoverageinformationNbCriticalextensionsC1
	Criticalextensionsfuture *UepagingcoverageinformationNbCriticalextensionsCriticalextensionsfuture
}
