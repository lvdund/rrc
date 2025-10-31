package ies

// SCGFailureInformationEUTRA-criticalExtensions ::= CHOICE
const (
	ScgfailureinformationeutraCriticalextensionsChoiceNothing = iota
	ScgfailureinformationeutraCriticalextensionsChoiceScgfailureinformationeutra
	ScgfailureinformationeutraCriticalextensionsChoiceCriticalextensionsfuture
)

type ScgfailureinformationeutraCriticalextensions struct {
	Choice                     uint64
	Scgfailureinformationeutra *Scgfailureinformationeutra
	Criticalextensionsfuture   *ScgfailureinformationeutraCriticalextensionsCriticalextensionsfuture
}
