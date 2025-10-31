package ies

// UEPagingCoverageInformation-criticalExtensions ::= CHOICE
const (
	UepagingcoverageinformationCriticalextensionsChoiceNothing = iota
	UepagingcoverageinformationCriticalextensionsChoiceC1
	UepagingcoverageinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type UepagingcoverageinformationCriticalextensions struct {
	Choice                   uint64
	C1                       *UepagingcoverageinformationCriticalextensionsC1
	Criticalextensionsfuture *UepagingcoverageinformationCriticalextensionsCriticalextensionsfuture
}
