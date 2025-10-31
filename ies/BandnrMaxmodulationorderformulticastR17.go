package ies

import "rrc/utils"

// BandNR-maxModulationOrderForMulticast-r17 ::= CHOICE
const (
	BandnrMaxmodulationorderformulticastR17ChoiceNothing = iota
	BandnrMaxmodulationorderformulticastR17ChoiceFr1R17
	BandnrMaxmodulationorderformulticastR17ChoiceFr2R17
)

type BandnrMaxmodulationorderformulticastR17 struct {
	Choice uint64
	Fr1R17 *utils.ENUMERATED
	Fr2R17 *utils.ENUMERATED
}
