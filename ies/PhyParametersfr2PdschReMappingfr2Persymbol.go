package ies

import "rrc/utils"

// Phy-ParametersFR2-pdsch-RE-MappingFR2-PerSymbol ::= ENUMERATED
type PhyParametersfr2PdschReMappingfr2Persymbol struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfr2PdschReMappingfr2PersymbolEnumeratedNothing = iota
	PhyParametersfr2PdschReMappingfr2PersymbolEnumeratedN6
	PhyParametersfr2PdschReMappingfr2PersymbolEnumeratedN20
)
