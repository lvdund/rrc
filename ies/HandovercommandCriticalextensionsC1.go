package ies

// HandoverCommand-criticalExtensions-c1 ::= CHOICE
const (
	HandovercommandCriticalextensionsC1ChoiceNothing = iota
	HandovercommandCriticalextensionsC1ChoiceHandovercommand
	HandovercommandCriticalextensionsC1ChoiceSpare3
	HandovercommandCriticalextensionsC1ChoiceSpare2
	HandovercommandCriticalextensionsC1ChoiceSpare1
)

type HandovercommandCriticalextensionsC1 struct {
	Choice          uint64
	Handovercommand *Handovercommand
	Spare3          *struct{}
	Spare2          *struct{}
	Spare1          *struct{}
}
