package ies

// ULInformationTransferMRDC-criticalExtensions-c1 ::= CHOICE
const (
	UlinformationtransfermrdcCriticalextensionsC1ChoiceNothing = iota
	UlinformationtransfermrdcCriticalextensionsC1ChoiceUlinformationtransfermrdc
	UlinformationtransfermrdcCriticalextensionsC1ChoiceSpare3
	UlinformationtransfermrdcCriticalextensionsC1ChoiceSpare2
	UlinformationtransfermrdcCriticalextensionsC1ChoiceSpare1
)

type UlinformationtransfermrdcCriticalextensionsC1 struct {
	Choice                    uint64
	Ulinformationtransfermrdc *Ulinformationtransfermrdc
	Spare3                    *struct{}
	Spare2                    *struct{}
	Spare1                    *struct{}
}
