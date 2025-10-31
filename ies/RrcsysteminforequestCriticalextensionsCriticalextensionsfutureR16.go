package ies

// RRCSystemInfoRequest-criticalExtensions-criticalExtensionsFuture-r16 ::= CHOICE
const (
	RrcsysteminforequestCriticalextensionsCriticalextensionsfutureR16ChoiceNothing = iota
	RrcsysteminforequestCriticalextensionsCriticalextensionsfutureR16ChoiceRrcpossysteminforequestR16
	RrcsysteminforequestCriticalextensionsCriticalextensionsfutureR16ChoiceCriticalextensionsfuture
)

type RrcsysteminforequestCriticalextensionsCriticalextensionsfutureR16 struct {
	Choice                     uint64
	RrcpossysteminforequestR16 *RrcPossysteminforequestR16
	Criticalextensionsfuture   *RrcsysteminforequestCriticalextensionsCriticalextensionsfutureR16Criticalextensionsfuture
}
