package ies

// MeasurementTimingConfiguration-criticalExtensions ::= CHOICE
const (
	MeasurementtimingconfigurationCriticalextensionsChoiceNothing = iota
	MeasurementtimingconfigurationCriticalextensionsChoiceC1
	MeasurementtimingconfigurationCriticalextensionsChoiceCriticalextensionsfuture
)

type MeasurementtimingconfigurationCriticalextensions struct {
	Choice                   uint64
	C1                       *MeasurementtimingconfigurationCriticalextensionsC1
	Criticalextensionsfuture *MeasurementtimingconfigurationCriticalextensionsCriticalextensionsfuture
}
