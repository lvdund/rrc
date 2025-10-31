package ies

// ULDedicatedMessageSegment-r16-criticalExtensions ::= CHOICE
const (
	UldedicatedmessagesegmentR16CriticalextensionsChoiceNothing = iota
	UldedicatedmessagesegmentR16CriticalextensionsChoiceUldedicatedmessagesegmentR16
	UldedicatedmessagesegmentR16CriticalextensionsChoiceCriticalextensionsfuture
)

type UldedicatedmessagesegmentR16Criticalextensions struct {
	Choice                       uint64
	UldedicatedmessagesegmentR16 *UldedicatedmessagesegmentR16
	Criticalextensionsfuture     *UldedicatedmessagesegmentR16CriticalextensionsCriticalextensionsfuture
}
