package ies

// BandCombinationList-UplinkTxSwitch-v1720 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1720
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1720 struct {
	Value []BandcombinationUplinktxswitchV1720 `lb:1,ub:maxBandComb`
}
