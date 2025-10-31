package ies

import "rrc/utils"

// BandNR-powerBoosting-pi2BPSK ::= ENUMERATED
type BandnrPowerboostingPi2bpsk struct {
	Value utils.ENUMERATED
}

const (
	BandnrPowerboostingPi2bpskEnumeratedNothing = iota
	BandnrPowerboostingPi2bpskEnumeratedSupported
)
