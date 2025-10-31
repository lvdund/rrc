package ies

// ParametersCDMA2000-r11-systemTimeInfo-r11 ::= CHOICE
const (
	Parameterscdma2000R11SystemtimeinfoR11ChoiceNothing = iota
	Parameterscdma2000R11SystemtimeinfoR11ChoiceExplicitvalue
	Parameterscdma2000R11SystemtimeinfoR11ChoiceDefaultvalue
)

type Parameterscdma2000R11SystemtimeinfoR11 struct {
	Choice        uint64
	Explicitvalue *Systemtimeinfocdma2000
	Defaultvalue  *struct{}
}
