package ies

import "rrc/utils"

// BandCombinationParameters-v1320 ::= SEQUENCE
type BandcombinationparametersV1320 struct {
	BandparameterlistV1320          *interface{}
	AdditionalrxTxPerformancereqR13 *utils.ENUMERATED
}
