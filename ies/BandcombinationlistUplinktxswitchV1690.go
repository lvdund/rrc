package ies

// BandCombinationList-UplinkTxSwitch-v1690 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1690
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1690 struct {
	Value []BandcombinationUplinktxswitchV1690 `lb:1,ub:maxBandComb`
}
