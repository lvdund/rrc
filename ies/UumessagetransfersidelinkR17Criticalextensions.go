package ies

// UuMessageTransferSidelink-r17-criticalExtensions ::= CHOICE
const (
	UumessagetransfersidelinkR17CriticalextensionsChoiceNothing = iota
	UumessagetransfersidelinkR17CriticalextensionsChoiceUumessagetransfersidelinkR17
	UumessagetransfersidelinkR17CriticalextensionsChoiceCriticalextensionsfuture
)

type UumessagetransfersidelinkR17Criticalextensions struct {
	Choice                       uint64
	UumessagetransfersidelinkR17 *UumessagetransfersidelinkR17
	Criticalextensionsfuture     *UumessagetransfersidelinkR17CriticalextensionsCriticalextensionsfuture
}
