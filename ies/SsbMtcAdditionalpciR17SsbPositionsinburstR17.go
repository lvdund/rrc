package ies

import "rrc/utils"

// SSB-MTC-AdditionalPCI-r17-ssb-PositionsInBurst-r17 ::= CHOICE
const (
	SsbMtcAdditionalpciR17SsbPositionsinburstR17ChoiceNothing = iota
	SsbMtcAdditionalpciR17SsbPositionsinburstR17ChoiceShortbitmap
	SsbMtcAdditionalpciR17SsbPositionsinburstR17ChoiceMediumbitmap
	SsbMtcAdditionalpciR17SsbPositionsinburstR17ChoiceLongbitmap
)

type SsbMtcAdditionalpciR17SsbPositionsinburstR17 struct {
	Choice       uint64
	Shortbitmap  *utils.BITSTRING `lb:4,ub:4`
	Mediumbitmap *utils.BITSTRING `lb:8,ub:8`
	Longbitmap   *utils.BITSTRING `lb:64,ub:64`
}
