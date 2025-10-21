package ies

import "rrc/utils"

// ParametersCDMA2000-r11 ::= SEQUENCE
// Extensible
type Parameterscdma2000R11 struct {
	SystemtimeinfoR11   *interface{}
	SearchwindowsizeR11 utils.INTEGER
	ParametershrpdR11   *interface{}
	Parameters1xrttR11  *interface{}
}
