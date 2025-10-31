package ies

// InDeviceCoexIndication-r11-criticalExtensions-c1 ::= CHOICE
const (
	IndevicecoexindicationR11CriticalextensionsC1ChoiceNothing = iota
	IndevicecoexindicationR11CriticalextensionsC1ChoiceIndevicecoexindicationR11
	IndevicecoexindicationR11CriticalextensionsC1ChoiceSpare3
	IndevicecoexindicationR11CriticalextensionsC1ChoiceSpare2
	IndevicecoexindicationR11CriticalextensionsC1ChoiceSpare1
)

type IndevicecoexindicationR11CriticalextensionsC1 struct {
	Choice                    uint64
	IndevicecoexindicationR11 *IndevicecoexindicationR11
	Spare3                    *struct{}
	Spare2                    *struct{}
	Spare1                    *struct{}
}
