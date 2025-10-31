package ies

// InterFreqRSTDMeasurementIndication-r10-criticalExtensions-c1 ::= CHOICE
const (
	InterfreqrstdmeasurementindicationR10CriticalextensionsC1ChoiceNothing = iota
	InterfreqrstdmeasurementindicationR10CriticalextensionsC1ChoiceInterfreqrstdmeasurementindicationR10
	InterfreqrstdmeasurementindicationR10CriticalextensionsC1ChoiceSpare3
	InterfreqrstdmeasurementindicationR10CriticalextensionsC1ChoiceSpare2
	InterfreqrstdmeasurementindicationR10CriticalextensionsC1ChoiceSpare1
)

type InterfreqrstdmeasurementindicationR10CriticalextensionsC1 struct {
	Choice                                uint64
	InterfreqrstdmeasurementindicationR10 *InterfreqrstdmeasurementindicationR10
	Spare3                                *struct{}
	Spare2                                *struct{}
	Spare1                                *struct{}
}
