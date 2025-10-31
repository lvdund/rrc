package ies

import "rrc/utils"

// RF-Parameters-NB-r13-multiNS-Pmax-r13 ::= ENUMERATED
type RfParametersNbR13MultinsPmaxR13 struct {
	Value utils.ENUMERATED
}

const (
	RfParametersNbR13MultinsPmaxR13EnumeratedNothing = iota
	RfParametersNbR13MultinsPmaxR13EnumeratedSupported
)
