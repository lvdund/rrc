package ies

// BandCombinationList-UplinkTxSwitch-v1640 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1640
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1640 struct {
	Value []BandcombinationUplinktxswitchV1640 `lb:1,ub:maxBandComb`
}
