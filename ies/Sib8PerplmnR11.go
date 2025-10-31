package ies

import "rrc/utils"

// SIB8-PerPLMN-r11 ::= SEQUENCE
type Sib8PerplmnR11 struct {
	PlmnIdentityR11       utils.INTEGER `lb:0,ub:maxPLMNR11`
	Parameterscdma2000R11 Sib8PerplmnR11Parameterscdma2000R11
}
