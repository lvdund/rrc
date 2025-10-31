package ies

// MeasurementReportSidelink-criticalExtensions ::= CHOICE
const (
	MeasurementreportsidelinkCriticalextensionsChoiceNothing = iota
	MeasurementreportsidelinkCriticalextensionsChoiceMeasurementreportsidelinkR16
	MeasurementreportsidelinkCriticalextensionsChoiceCriticalextensionsfuture
)

type MeasurementreportsidelinkCriticalextensions struct {
	Choice                       uint64
	MeasurementreportsidelinkR16 *MeasurementreportsidelinkR16
	Criticalextensionsfuture     *MeasurementreportsidelinkCriticalextensionsCriticalextensionsfuture
}
