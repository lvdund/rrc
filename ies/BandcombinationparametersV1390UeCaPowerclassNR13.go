package ies

import "rrc/utils"

// BandCombinationParameters-v1390-ue-CA-PowerClass-N-r13 ::= ENUMERATED
type BandcombinationparametersV1390UeCaPowerclassNR13 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1390UeCaPowerclassNR13EnumeratedNothing = iota
	BandcombinationparametersV1390UeCaPowerclassNR13EnumeratedClass2
)
