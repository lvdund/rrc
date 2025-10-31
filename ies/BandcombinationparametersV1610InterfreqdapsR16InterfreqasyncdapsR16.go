package ies

import "rrc/utils"

// BandCombinationParameters-v1610-interFreqDAPS-r16-interFreqAsyncDAPS-r16 ::= ENUMERATED
type BandcombinationparametersV1610InterfreqdapsR16InterfreqasyncdapsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1610InterfreqdapsR16InterfreqasyncdapsR16EnumeratedNothing = iota
	BandcombinationparametersV1610InterfreqdapsR16InterfreqasyncdapsR16EnumeratedSupported
)
