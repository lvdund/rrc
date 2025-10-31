package ies

import "rrc/utils"

// ParametersCDMA2000-r11 ::= SEQUENCE
// Extensible
type Parameterscdma2000R11 struct {
	SystemtimeinfoR11   *Parameterscdma2000R11SystemtimeinfoR11
	SearchwindowsizeR11 utils.INTEGER `lb:0,ub:15`
	ParametershrpdR11   *Parameterscdma2000R11ParametershrpdR11
	Parameters1xrttR11  *Parameterscdma2000R11Parameters1xrttR11
}
