package ies

// MBMSInterestIndication-r11-criticalExtensions-c1 ::= CHOICE
const (
	MbmsinterestindicationR11CriticalextensionsC1ChoiceNothing = iota
	MbmsinterestindicationR11CriticalextensionsC1ChoiceInterestindicationR11
	MbmsinterestindicationR11CriticalextensionsC1ChoiceSpare3
	MbmsinterestindicationR11CriticalextensionsC1ChoiceSpare2
	MbmsinterestindicationR11CriticalextensionsC1ChoiceSpare1
)

type MbmsinterestindicationR11CriticalextensionsC1 struct {
	Choice                uint64
	InterestindicationR11 *MbmsinterestindicationR11
	Spare3                *struct{}
	Spare2                *struct{}
	Spare1                *struct{}
}
