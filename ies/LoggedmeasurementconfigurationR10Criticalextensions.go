package ies

// LoggedMeasurementConfiguration-r10-criticalExtensions ::= CHOICE
const (
	LoggedmeasurementconfigurationR10CriticalextensionsChoiceNothing = iota
	LoggedmeasurementconfigurationR10CriticalextensionsChoiceC1
	LoggedmeasurementconfigurationR10CriticalextensionsChoiceCriticalextensionsfuture
)

type LoggedmeasurementconfigurationR10Criticalextensions struct {
	Choice                   uint64
	C1                       *LoggedmeasurementconfigurationR10CriticalextensionsC1
	Criticalextensionsfuture *LoggedmeasurementconfigurationR10CriticalextensionsCriticalextensionsfuture
}
