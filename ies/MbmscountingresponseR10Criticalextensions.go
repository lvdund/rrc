package ies

// MBMSCountingResponse-r10-criticalExtensions ::= CHOICE
const (
	MbmscountingresponseR10CriticalextensionsChoiceNothing = iota
	MbmscountingresponseR10CriticalextensionsChoiceC1
	MbmscountingresponseR10CriticalextensionsChoiceCriticalextensionsfuture
)

type MbmscountingresponseR10Criticalextensions struct {
	Choice                   uint64
	C1                       *MbmscountingresponseR10CriticalextensionsC1
	Criticalextensionsfuture *MbmscountingresponseR10CriticalextensionsCriticalextensionsfuture
}
