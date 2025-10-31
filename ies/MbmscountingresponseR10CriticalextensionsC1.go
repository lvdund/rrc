package ies

// MBMSCountingResponse-r10-criticalExtensions-c1 ::= CHOICE
const (
	MbmscountingresponseR10CriticalextensionsC1ChoiceNothing = iota
	MbmscountingresponseR10CriticalextensionsC1ChoiceCountingresponseR10
	MbmscountingresponseR10CriticalextensionsC1ChoiceSpare3
	MbmscountingresponseR10CriticalextensionsC1ChoiceSpare2
	MbmscountingresponseR10CriticalextensionsC1ChoiceSpare1
)

type MbmscountingresponseR10CriticalextensionsC1 struct {
	Choice              uint64
	CountingresponseR10 *MbmscountingresponseR10
	Spare3              *struct{}
	Spare2              *struct{}
	Spare1              *struct{}
}
