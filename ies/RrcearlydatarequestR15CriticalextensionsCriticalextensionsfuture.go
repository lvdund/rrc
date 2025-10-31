package ies

// RRCEarlyDataRequest-r15-criticalExtensions-criticalExtensionsFuture ::= CHOICE
const (
	RrcearlydatarequestR15CriticalextensionsCriticalextensionsfutureChoiceNothing = iota
	RrcearlydatarequestR15CriticalextensionsCriticalextensionsfutureChoiceRrcearlydatarequest5gcR16
	RrcearlydatarequestR15CriticalextensionsCriticalextensionsfutureChoiceCriticalextensionsfutureR16
)

type RrcearlydatarequestR15CriticalextensionsCriticalextensionsfuture struct {
	Choice                      uint64
	Rrcearlydatarequest5gcR16   *Rrcearlydatarequest5gcR16
	CriticalextensionsfutureR16 *RrcearlydatarequestR15CriticalextensionsCriticalextensionsfutureCriticalextensionsfutureR16
}
