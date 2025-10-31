package ies

// RRCSystemInfoRequest-criticalExtensions ::= CHOICE
const (
	RrcsysteminforequestCriticalextensionsChoiceNothing = iota
	RrcsysteminforequestCriticalextensionsChoiceRrcsysteminforequest
	RrcsysteminforequestCriticalextensionsChoiceCriticalextensionsfutureR16
)

type RrcsysteminforequestCriticalextensions struct {
	Choice                      uint64
	Rrcsysteminforequest        *Rrcsysteminforequest
	CriticalextensionsfutureR16 *RrcsysteminforequestCriticalextensionsCriticalextensionsfutureR16
}
