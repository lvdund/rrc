package ies

import "rrc/utils"

// MIMO-UE-Parameters-r13 ::= SEQUENCE
type MimoUeParametersR13 struct {
	Parameterstm9R13               *MimoUeParameterspertmR13
	Parameterstm10R13              *MimoUeParameterspertmR13
	SrsEnhancementstddR13          *utils.ENUMERATED
	SrsEnhancementsR13             *utils.ENUMERATED
	InterferencemeasrestrictionR13 *utils.ENUMERATED
}
