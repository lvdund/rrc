package ies

// MeasurementReportAppLayer-r17-criticalExtensions ::= CHOICE
const (
	MeasurementreportapplayerR17CriticalextensionsChoiceNothing = iota
	MeasurementreportapplayerR17CriticalextensionsChoiceMeasurementreportapplayerR17
	MeasurementreportapplayerR17CriticalextensionsChoiceCriticalextensionsfuture
)

type MeasurementreportapplayerR17Criticalextensions struct {
	Choice                       uint64
	MeasurementreportapplayerR17 *MeasurementreportapplayerR17
	Criticalextensionsfuture     *MeasurementreportapplayerR17CriticalextensionsCriticalextensionsfuture
}
