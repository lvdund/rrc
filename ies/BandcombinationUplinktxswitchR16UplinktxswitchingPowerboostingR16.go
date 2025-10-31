package ies

import "rrc/utils"

// BandCombination-UplinkTxSwitch-r16-uplinkTxSwitching-PowerBoosting-r16 ::= ENUMERATED
type BandcombinationUplinktxswitchR16UplinktxswitchingPowerboostingR16 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationUplinktxswitchR16UplinktxswitchingPowerboostingR16EnumeratedNothing = iota
	BandcombinationUplinktxswitchR16UplinktxswitchingPowerboostingR16EnumeratedSupported
)
