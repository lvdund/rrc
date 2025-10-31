package ies

// BandCombinationList-UplinkTxSwitch-r16 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-r16
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchR16 struct {
	Value []BandcombinationUplinktxswitchR16 `lb:1,ub:maxBandComb`
}
