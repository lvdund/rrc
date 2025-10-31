package ies

import "rrc/utils"

// ULTxSwitchingBandPair-r16-uplinkTxSwitchingPeriod-r16 ::= ENUMERATED
type UltxswitchingbandpairR16UplinktxswitchingperiodR16 struct {
	Value utils.ENUMERATED
}

const (
	UltxswitchingbandpairR16UplinktxswitchingperiodR16EnumeratedNothing = iota
	UltxswitchingbandpairR16UplinktxswitchingperiodR16EnumeratedN35us
	UltxswitchingbandpairR16UplinktxswitchingperiodR16EnumeratedN140us
	UltxswitchingbandpairR16UplinktxswitchingperiodR16EnumeratedN210us
)
