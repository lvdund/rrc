package ies

// SecurityModeCommand-criticalExtensions-c1 ::= CHOICE
const (
	SecuritymodecommandCriticalextensionsC1ChoiceNothing = iota
	SecuritymodecommandCriticalextensionsC1ChoiceSecuritymodecommandR8
	SecuritymodecommandCriticalextensionsC1ChoiceSpare3
	SecuritymodecommandCriticalextensionsC1ChoiceSpare2
	SecuritymodecommandCriticalextensionsC1ChoiceSpare1
)

type SecuritymodecommandCriticalextensionsC1 struct {
	Choice                uint64
	SecuritymodecommandR8 *SecuritymodecommandR8
	Spare3                *struct{}
	Spare2                *struct{}
	Spare1                *struct{}
}
