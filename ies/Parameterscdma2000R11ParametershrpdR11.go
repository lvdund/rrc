package ies

// ParametersCDMA2000-r11-parametersHRPD-r11 ::= SEQUENCE
type Parameterscdma2000R11ParametershrpdR11 struct {
	PreregistrationinfohrpdR11       Preregistrationinfohrpd
	CellreselectionparametershrpdR11 *Cellreselectionparameterscdma2000R11
}
