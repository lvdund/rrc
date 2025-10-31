package ies

// DLDedicatedMessageSegment-r16-criticalExtensions ::= CHOICE
const (
	DldedicatedmessagesegmentR16CriticalextensionsChoiceNothing = iota
	DldedicatedmessagesegmentR16CriticalextensionsChoiceDldedicatedmessagesegmentR16
	DldedicatedmessagesegmentR16CriticalextensionsChoiceCriticalextensionsfuture
)

type DldedicatedmessagesegmentR16Criticalextensions struct {
	Choice                       uint64
	DldedicatedmessagesegmentR16 *DldedicatedmessagesegmentR16
	Criticalextensionsfuture     *DldedicatedmessagesegmentR16CriticalextensionsCriticalextensionsfuture
}
