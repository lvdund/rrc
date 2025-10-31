package ies

// CG-Config-criticalExtensions ::= CHOICE
const (
	CgConfigCriticalextensionsChoiceNothing = iota
	CgConfigCriticalextensionsChoiceC1
	CgConfigCriticalextensionsChoiceCriticalextensionsfuture
)

type CgConfigCriticalextensions struct {
	Choice                   uint64
	C1                       *CgConfigCriticalextensionsC1
	Criticalextensionsfuture *CgConfigCriticalextensionsCriticalextensionsfuture
}
