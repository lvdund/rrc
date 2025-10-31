package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-RetuningSymbols-r14 ::= ENUMERATED
type PhylayerparametersV1430CeRetuningsymbolsR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CeRetuningsymbolsR14EnumeratedNothing = iota
	PhylayerparametersV1430CeRetuningsymbolsR14EnumeratedN0
	PhylayerparametersV1430CeRetuningsymbolsR14EnumeratedN1
)
