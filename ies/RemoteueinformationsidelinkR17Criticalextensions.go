package ies

// RemoteUEInformationSidelink-r17-criticalExtensions ::= CHOICE
const (
	RemoteueinformationsidelinkR17CriticalextensionsChoiceNothing = iota
	RemoteueinformationsidelinkR17CriticalextensionsChoiceRemoteueinformationsidelinkR17
	RemoteueinformationsidelinkR17CriticalextensionsChoiceCriticalextensionsfuture
)

type RemoteueinformationsidelinkR17Criticalextensions struct {
	Choice                         uint64
	RemoteueinformationsidelinkR17 *RemoteueinformationsidelinkR17
	Criticalextensionsfuture       *RemoteueinformationsidelinkR17CriticalextensionsCriticalextensionsfuture
}
