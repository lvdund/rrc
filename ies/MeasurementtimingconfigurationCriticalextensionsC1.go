package ies

// MeasurementTimingConfiguration-criticalExtensions-c1 ::= CHOICE
const (
	MeasurementtimingconfigurationCriticalextensionsC1ChoiceNothing = iota
	MeasurementtimingconfigurationCriticalextensionsC1ChoiceMeastimingconf
	MeasurementtimingconfigurationCriticalextensionsC1ChoiceSpare3
	MeasurementtimingconfigurationCriticalextensionsC1ChoiceSpare2
	MeasurementtimingconfigurationCriticalextensionsC1ChoiceSpare1
)

type MeasurementtimingconfigurationCriticalextensionsC1 struct {
	Choice         uint64
	Meastimingconf *Measurementtimingconfiguration
	Spare3         *struct{}
	Spare2         *struct{}
	Spare1         *struct{}
}
