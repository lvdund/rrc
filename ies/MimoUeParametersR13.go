package ies

// MIMO-UE-Parameters-r13 ::= SEQUENCE
type MimoUeParametersR13 struct {
	Parameterstm9R13               *MimoUeParameterspertmR13
	Parameterstm10R13              *MimoUeParameterspertmR13
	SrsEnhancementstddR13          *MimoUeParametersR13SrsEnhancementstddR13
	SrsEnhancementsR13             *MimoUeParametersR13SrsEnhancementsR13
	InterferencemeasrestrictionR13 *MimoUeParametersR13InterferencemeasrestrictionR13
}
