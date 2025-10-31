package ies

// BandCombinationList-UplinkTxSwitch-v1700 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1700
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1700 struct {
	Value []BandcombinationUplinktxswitchV1700 `lb:1,ub:maxBandComb`
}
