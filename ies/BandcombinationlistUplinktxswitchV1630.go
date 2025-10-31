package ies

// BandCombinationList-UplinkTxSwitch-v1630 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1630
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1630 struct {
	Value []BandcombinationUplinktxswitchV1630 `lb:1,ub:maxBandComb`
}
