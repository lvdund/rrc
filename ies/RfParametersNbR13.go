package ies

// RF-Parameters-NB-r13 ::= SEQUENCE
type RfParametersNbR13 struct {
	SupportedbandlistR13 SupportedbandlistNbR13
	MultinsPmaxR13       *RfParametersNbR13MultinsPmaxR13
}
