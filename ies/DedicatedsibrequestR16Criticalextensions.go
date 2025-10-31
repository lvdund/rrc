package ies

// DedicatedSIBRequest-r16-criticalExtensions ::= CHOICE
const (
	DedicatedsibrequestR16CriticalextensionsChoiceNothing = iota
	DedicatedsibrequestR16CriticalextensionsChoiceDedicatedsibrequestR16
	DedicatedsibrequestR16CriticalextensionsChoiceCriticalextensionsfuture
)

type DedicatedsibrequestR16Criticalextensions struct {
	Choice                   uint64
	DedicatedsibrequestR16   *DedicatedsibrequestR16
	Criticalextensionsfuture *DedicatedsibrequestR16CriticalextensionsCriticalextensionsfuture
}
