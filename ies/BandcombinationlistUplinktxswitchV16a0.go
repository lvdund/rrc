package ies

// BandCombinationList-UplinkTxSwitch-v16a0 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v16a0
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV16a0 struct {
	Value []BandcombinationUplinktxswitchV16a0 `lb:1,ub:maxBandComb`
}
