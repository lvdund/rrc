package ies

// InDeviceCoexIndication-r11-criticalExtensions ::= CHOICE
const (
	IndevicecoexindicationR11CriticalextensionsChoiceNothing = iota
	IndevicecoexindicationR11CriticalextensionsChoiceC1
	IndevicecoexindicationR11CriticalextensionsChoiceCriticalextensionsfuture
)

type IndevicecoexindicationR11Criticalextensions struct {
	Choice                   uint64
	C1                       *IndevicecoexindicationR11CriticalextensionsC1
	Criticalextensionsfuture *IndevicecoexindicationR11CriticalextensionsCriticalextensionsfuture
}
