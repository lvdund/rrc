package ies

// BandCombinationList-UplinkTxSwitch-v1650 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1650
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1650 struct {
	Value []BandcombinationUplinktxswitchV1650 `lb:1,ub:maxBandComb`
}
