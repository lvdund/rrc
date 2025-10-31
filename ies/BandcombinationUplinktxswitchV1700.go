package ies

// BandCombination-UplinkTxSwitch-v1700 ::= SEQUENCE
type BandcombinationUplinktxswitchV1700 struct {
	BandcombinationV1700                     *BandcombinationV1700
	SupportedbandpairlistnrV1700             *[]UltxswitchingbandpairV1700           `lb:1,ub:maxULTxSwitchingBandPairs`
	UplinktxswitchingbandparameterslistV1700 *[]UplinktxswitchingbandparametersV1700 `lb:1,ub:maxSimultaneousBands`
}
