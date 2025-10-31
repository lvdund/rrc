package ies

// MeasurementReport-criticalExtensions ::= CHOICE
const (
	MeasurementreportCriticalextensionsChoiceNothing = iota
	MeasurementreportCriticalextensionsChoiceMeasurementreport
	MeasurementreportCriticalextensionsChoiceCriticalextensionsfuture
)

type MeasurementreportCriticalextensions struct {
	Choice                   uint64
	Measurementreport        *Measurementreport
	Criticalextensionsfuture *MeasurementreportCriticalextensionsCriticalextensionsfuture
}
