package ies

// MeasurementReport-criticalExtensions ::= CHOICE
const (
	MeasurementreportCriticalextensionsChoiceNothing = iota
	MeasurementreportCriticalextensionsChoiceC1
	MeasurementreportCriticalextensionsChoiceCriticalextensionsfuture
)

type MeasurementreportCriticalextensions struct {
	Choice                   uint64
	C1                       *MeasurementreportCriticalextensionsC1
	Criticalextensionsfuture *MeasurementreportCriticalextensionsCriticalextensionsfuture
}
