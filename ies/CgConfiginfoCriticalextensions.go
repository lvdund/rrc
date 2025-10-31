package ies

// CG-ConfigInfo-criticalExtensions ::= CHOICE
const (
	CgConfiginfoCriticalextensionsChoiceNothing = iota
	CgConfiginfoCriticalextensionsChoiceC1
	CgConfiginfoCriticalextensionsChoiceCriticalextensionsfuture
)

type CgConfiginfoCriticalextensions struct {
	Choice                   uint64
	C1                       *CgConfiginfoCriticalextensionsC1
	Criticalextensionsfuture *CgConfiginfoCriticalextensionsCriticalextensionsfuture
}
