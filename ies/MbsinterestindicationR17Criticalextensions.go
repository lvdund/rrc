package ies

// MBSInterestIndication-r17-criticalExtensions ::= CHOICE
const (
	MbsinterestindicationR17CriticalextensionsChoiceNothing = iota
	MbsinterestindicationR17CriticalextensionsChoiceMbsinterestindicationR17
	MbsinterestindicationR17CriticalextensionsChoiceCriticalextensionsfuture
)

type MbsinterestindicationR17Criticalextensions struct {
	Choice                   uint64
	MbsinterestindicationR17 *MbsinterestindicationR17
	Criticalextensionsfuture *MbsinterestindicationR17CriticalextensionsCriticalextensionsfuture
}
