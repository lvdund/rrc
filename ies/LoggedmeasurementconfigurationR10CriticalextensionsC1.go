package ies

// LoggedMeasurementConfiguration-r10-criticalExtensions-c1 ::= CHOICE
const (
	LoggedmeasurementconfigurationR10CriticalextensionsC1ChoiceNothing = iota
	LoggedmeasurementconfigurationR10CriticalextensionsC1ChoiceLoggedmeasurementconfigurationR10
	LoggedmeasurementconfigurationR10CriticalextensionsC1ChoiceSpare3
	LoggedmeasurementconfigurationR10CriticalextensionsC1ChoiceSpare2
	LoggedmeasurementconfigurationR10CriticalextensionsC1ChoiceSpare1
)

type LoggedmeasurementconfigurationR10CriticalextensionsC1 struct {
	Choice                            uint64
	LoggedmeasurementconfigurationR10 *LoggedmeasurementconfigurationR10
	Spare3                            *struct{}
	Spare2                            *struct{}
	Spare1                            *struct{}
}
