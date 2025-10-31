package ies

// WLANConnectionStatusReport-r13-criticalExtensions ::= CHOICE
const (
	WlanconnectionstatusreportR13CriticalextensionsChoiceNothing = iota
	WlanconnectionstatusreportR13CriticalextensionsChoiceC1
	WlanconnectionstatusreportR13CriticalextensionsChoiceCriticalextensionsfuture
)

type WlanconnectionstatusreportR13Criticalextensions struct {
	Choice                   uint64
	C1                       *WlanconnectionstatusreportR13CriticalextensionsC1
	Criticalextensionsfuture *WlanconnectionstatusreportR13CriticalextensionsCriticalextensionsfuture
}
