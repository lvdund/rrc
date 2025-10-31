package ies

import "rrc/utils"

// Phy-ParametersFR1-pdsch-RE-MappingFR1-PerSymbol ::= ENUMERATED
type PhyParametersfr1PdschReMappingfr1Persymbol struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr1PdschReMappingfr1PersymbolEnumeratedNothing = iota
	PhyParametersfr1PdschReMappingfr1PersymbolEnumeratedN10
	PhyParametersfr1PdschReMappingfr1PersymbolEnumeratedN20
)
