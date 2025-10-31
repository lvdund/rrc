package ies

// BandCombinationList-UplinkTxSwitch-v1670 ::= SEQUENCE OF BandCombination-UplinkTxSwitch-v1670
// SIZE (1..maxBandComb)
type BandcombinationlistUplinktxswitchV1670 struct {
	Value []BandcombinationUplinktxswitchV1670 `lb:1,ub:maxBandComb`
}
