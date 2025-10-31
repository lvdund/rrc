package ies

// MBMSInterestIndication-r11-criticalExtensions ::= CHOICE
const (
	MbmsinterestindicationR11CriticalextensionsChoiceNothing = iota
	MbmsinterestindicationR11CriticalextensionsChoiceC1
	MbmsinterestindicationR11CriticalextensionsChoiceCriticalextensionsfuture
)

type MbmsinterestindicationR11Criticalextensions struct {
	Choice                   uint64
	C1                       *MbmsinterestindicationR11CriticalextensionsC1
	Criticalextensionsfuture *MbmsinterestindicationR11CriticalextensionsCriticalextensionsfuture
}
