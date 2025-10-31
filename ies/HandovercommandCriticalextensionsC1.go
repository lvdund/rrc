package ies

// HandoverCommand-criticalExtensions-c1 ::= CHOICE
const (
	HandovercommandCriticalextensionsC1ChoiceNothing = iota
	HandovercommandCriticalextensionsC1ChoiceHandovercommandR8
	HandovercommandCriticalextensionsC1ChoiceSpare7
	HandovercommandCriticalextensionsC1ChoiceSpare6
	HandovercommandCriticalextensionsC1ChoiceSpare5
	HandovercommandCriticalextensionsC1ChoiceSpare4
	HandovercommandCriticalextensionsC1ChoiceSpare3
	HandovercommandCriticalextensionsC1ChoiceSpare2
	HandovercommandCriticalextensionsC1ChoiceSpare1
)

type HandovercommandCriticalextensionsC1 struct {
	Choice            uint64
	HandovercommandR8 *HandovercommandR8
	Spare7            *struct{}
	Spare6            *struct{}
	Spare5            *struct{}
	Spare4            *struct{}
	Spare3            *struct{}
	Spare2            *struct{}
	Spare1            *struct{}
}
