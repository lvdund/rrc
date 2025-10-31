package ies

// BandCombinationList-UplinkTxSwitch-v1730 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1730
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1730 struct {
	Value []BandcombinationUplinktxswitchV1730 `lb:1,ub:maxBandComb`
}
