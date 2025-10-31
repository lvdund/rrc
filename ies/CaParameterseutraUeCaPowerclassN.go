package ies

import "rrc/utils"

// CA-ParametersEUTRA-ue-CA-PowerClass-N ::= ENUMERATED
type CaParameterseutraUeCaPowerclassN struct {
	Value utils.ENUMERATED
}

const (
	CaParameterseutraUeCaPowerclassNEnumeratedNothing = iota
	CaParameterseutraUeCaPowerclassNEnumeratedClass2
)
