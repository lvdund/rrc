package ies

// SIB8-PerPLMN-r11-parametersCDMA2000-r11 ::= CHOICE
const (
	Sib8PerplmnR11Parameterscdma2000R11ChoiceNothing = iota
	Sib8PerplmnR11Parameterscdma2000R11ChoiceExplicitvalue
	Sib8PerplmnR11Parameterscdma2000R11ChoiceDefaultvalue
)

type Sib8PerplmnR11Parameterscdma2000R11 struct {
	Choice        uint64
	Explicitvalue *Parameterscdma2000R11
	Defaultvalue  *struct{}
}
