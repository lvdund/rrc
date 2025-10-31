package ies

// MeasurementReport-criticalExtensions-c1 ::= CHOICE
const (
	MeasurementreportCriticalextensionsC1ChoiceNothing = iota
	MeasurementreportCriticalextensionsC1ChoiceMeasurementreportR8
	MeasurementreportCriticalextensionsC1ChoiceSpare7
	MeasurementreportCriticalextensionsC1ChoiceSpare6
	MeasurementreportCriticalextensionsC1ChoiceSpare5
	MeasurementreportCriticalextensionsC1ChoiceSpare4
	MeasurementreportCriticalextensionsC1ChoiceSpare3
	MeasurementreportCriticalextensionsC1ChoiceSpare2
	MeasurementreportCriticalextensionsC1ChoiceSpare1
)

type MeasurementreportCriticalextensionsC1 struct {
	Choice              uint64
	MeasurementreportR8 *MeasurementreportR8
	Spare7              *struct{}
	Spare6              *struct{}
	Spare5              *struct{}
	Spare4              *struct{}
	Spare3              *struct{}
	Spare2              *struct{}
	Spare1              *struct{}
}
