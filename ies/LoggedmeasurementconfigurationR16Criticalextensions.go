package ies

// LoggedMeasurementConfiguration-r16-criticalExtensions ::= CHOICE
const (
	LoggedmeasurementconfigurationR16CriticalextensionsChoiceNothing = iota
	LoggedmeasurementconfigurationR16CriticalextensionsChoiceLoggedmeasurementconfigurationR16
	LoggedmeasurementconfigurationR16CriticalextensionsChoiceCriticalextensionsfuture
)

type LoggedmeasurementconfigurationR16Criticalextensions struct {
	Choice                            uint64
	LoggedmeasurementconfigurationR16 *LoggedmeasurementconfigurationR16
	Criticalextensionsfuture          *LoggedmeasurementconfigurationR16CriticalextensionsCriticalextensionsfuture
}
