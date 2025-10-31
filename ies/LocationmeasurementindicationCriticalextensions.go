package ies

// LocationMeasurementIndication-criticalExtensions ::= CHOICE
const (
	LocationmeasurementindicationCriticalextensionsChoiceNothing = iota
	LocationmeasurementindicationCriticalextensionsChoiceLocationmeasurementindication
	LocationmeasurementindicationCriticalextensionsChoiceCriticalextensionsfuture
)

type LocationmeasurementindicationCriticalextensions struct {
	Choice                        uint64
	Locationmeasurementindication *Locationmeasurementindication
	Criticalextensionsfuture      *LocationmeasurementindicationCriticalextensionsCriticalextensionsfuture
}
