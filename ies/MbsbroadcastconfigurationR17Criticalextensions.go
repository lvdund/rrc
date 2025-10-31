package ies

// MBSBroadcastConfiguration-r17-criticalExtensions ::= CHOICE
const (
	MbsbroadcastconfigurationR17CriticalextensionsChoiceNothing = iota
	MbsbroadcastconfigurationR17CriticalextensionsChoiceMbsbroadcastconfigurationR17
	MbsbroadcastconfigurationR17CriticalextensionsChoiceCriticalextensionsfuture
)

type MbsbroadcastconfigurationR17Criticalextensions struct {
	Choice                       uint64
	MbsbroadcastconfigurationR17 *MbsbroadcastconfigurationR17
	Criticalextensionsfuture     *MbsbroadcastconfigurationR17CriticalextensionsCriticalextensionsfuture
}
