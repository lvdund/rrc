package ies

// MeasReportAppLayer-r15-criticalExtensions ::= CHOICE
const (
	MeasreportapplayerR15CriticalextensionsChoiceNothing = iota
	MeasreportapplayerR15CriticalextensionsChoiceMeasreportapplayerR15
	MeasreportapplayerR15CriticalextensionsChoiceCriticalextensionsfuture
)

type MeasreportapplayerR15Criticalextensions struct {
	Choice                   uint64
	MeasreportapplayerR15    *MeasreportapplayerR15
	Criticalextensionsfuture *MeasreportapplayerR15CriticalextensionsCriticalextensionsfuture
}
