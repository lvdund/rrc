package ies

// InterFreqRSTDMeasurementIndication-r10-criticalExtensions ::= CHOICE
const (
	InterfreqrstdmeasurementindicationR10CriticalextensionsChoiceNothing = iota
	InterfreqrstdmeasurementindicationR10CriticalextensionsChoiceC1
	InterfreqrstdmeasurementindicationR10CriticalextensionsChoiceCriticalextensionsfuture
)

type InterfreqrstdmeasurementindicationR10Criticalextensions struct {
	Choice                   uint64
	C1                       *InterfreqrstdmeasurementindicationR10CriticalextensionsC1
	Criticalextensionsfuture *InterfreqrstdmeasurementindicationR10CriticalextensionsCriticalextensionsfuture
}
