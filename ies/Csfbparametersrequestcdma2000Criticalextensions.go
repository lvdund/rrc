package ies

// CSFBParametersRequestCDMA2000-criticalExtensions ::= CHOICE
const (
	Csfbparametersrequestcdma2000CriticalextensionsChoiceNothing = iota
	Csfbparametersrequestcdma2000CriticalextensionsChoiceCsfbparametersrequestcdma2000R8
	Csfbparametersrequestcdma2000CriticalextensionsChoiceCriticalextensionsfuture
)

type Csfbparametersrequestcdma2000Criticalextensions struct {
	Choice                          uint64
	Csfbparametersrequestcdma2000R8 *Csfbparametersrequestcdma2000R8
	Criticalextensionsfuture        *Csfbparametersrequestcdma2000CriticalextensionsCriticalextensionsfuture
}
